package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os/exec"
	"strconv"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

type Builder struct {
	FromUser     string
	Secret       string
	ToUsers      string
	Subject      string
	Type         string
	Server       string
	Port         string
	Body         string
	EmailContent EmailContent
}

func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}
	if envs["FROM_USER"] == "" {
		return nil, fmt.Errorf("environment variable FROM_USER is requried")
	} else {
		b.FromUser = envs["FROM_USER"]
		//fmt.Println(b.FromUser)
	}
	if envs["TO_USERS"] == "" {
		return nil, fmt.Errorf("environment variable TO_USER is requried")
	} else {
		b.ToUsers = envs["TO_USERS"]
		//fmt.Println(b.ToUsers)
	}
	if envs["SECRET"] == "" {
		return nil, fmt.Errorf("environment variable SECRET is requried")
	} else {
		b.Secret = envs["SECRET"]
		//fmt.Println(b.Secret)
	}
	if envs["SMTP_SERVER_PORT"] == "" {
		return nil, fmt.Errorf("environment variable SMTP_SERVER_PORT is requried")
	} else {
		//fmt.Println(envs["SMTP_SERVER_PORT"])
		param := strings.SplitN(envs["SMTP_SERVER_PORT"], ":", 2)
		b.Server = param[0]
		b.Port = param[1]
		fmt.Printf("smtp_server: %s, smtp_port: %s\n", b.Server, b.Port)
	}

	b.Subject = envs["SUBJECT"]

	if envs["TEXT"] != "" {
		b.Type = "text"
		b.Body = envs["TEXT"]
		return b, nil
	}

	task := &FlowTask{}
	err := json.Unmarshal([]byte(envs["_WORKFLOW_TASK_DETAIL"]), task)
	if err != nil {
		return nil, err
	}
	var totalTime string
	if task.End != nil && task.Start != nil {
		totalTime = fmt.Sprintf("总耗时: %d 秒", (int64)(task.End.Sub(*task.Start).Seconds()))
	}
	fmt.Println(totalTime)

	emailContent := EmailContent{
		Title: fmt.Sprintf("%s/%s/%s", task.Namespace, task.Repo, task.Name),
		Text:  fmt.Sprintf("状态: %s\n%s", task.Status, totalTime),
		Url:   task.DetailURL,
	}
	b.EmailContent = emailContent
	b.Type = "html"
	return b, nil
}

func (b *Builder) run(templateName string) error {
	if b.Type == "html" {
		items := make(map[string]string)

		items["t_title"] = b.EmailContent.Title
		items["t_content"] = b.EmailContent.Text
		items["t_url"] = b.EmailContent.Url
		err := b.ParseTemplate(templateName, items)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := b.SendEmail(); err != nil {
		log.Printf("Failed to send the email to %s\n", b.ToUsers)
		return err
	} else {
		log.Printf("Email has been sent to %s\n", b.ToUsers)
	}
	return nil
}

func (b *Builder) ParseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err := t.Execute(buffer, data); err != nil {
		return err
	}
	b.Body = buffer.String()
	return nil
}

func (b *Builder) SendEmail() error {
	var toUsers = strings.Split(b.ToUsers, ",")
	m := gomail.NewMessage()
	//设置发件人
	m.SetAddressHeader("From", b.FromUser, "工作流邮件通知")
	//设置收件人
	m.SetHeader("To", toUsers...)
	//设置主题
	m.SetHeader("Subject", b.Subject)
	//设置正文
	//m.SetBody("text", "hello world!")

	m.SetBody("text/html", b.Body)
	//设置发送邮件服务器、端口、发件人账号、发件人密码
	port, err := strconv.Atoi(b.Port)
	if err != nil {
		return err
	}
	d := gomail.NewPlainDialer(b.Server, port, b.FromUser, b.Secret)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

type CMD struct {
	Command []string
	WorkDir string
}

func (c CMD) Run() (string, error) {
	fmt.Println("Run CMD: ", strings.Join(c.Command, "工作流通知"))

	cmd := exec.Command(c.Command[0], c.Command[1:]...)
	if c.WorkDir != "" {
		cmd.Dir = c.WorkDir
	}
	data, err := cmd.CombinedOutput()

	result := string(data)
	if len(result) > 0 {
		fmt.Println(result)

	}

	return result, err
}

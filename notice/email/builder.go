package main

import (
	"fmt"
	"os/exec"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

type Builder struct {
	FromUser string
	Secret   string
	ToUsers  string
	Subject  string
	Text     string
}

func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}
	if envs["FROM_USER"] != "" {
		b.FromUser = envs["FROM_USER"]
	} else {
		return nil, fmt.Errorf("environment variable FROM_USER is requried")
	}
	if envs["TO_USERS"] != "" {
		b.ToUsers = envs["TO_USERS"]
	} else {
		return nil, fmt.Errorf("environment variable TO_USER is requried")
	}
	if envs["SECRET"] != "" {
		b.Secret = envs["SECRET"]
	} else {
		return nil, fmt.Errorf("environment variable SECRET is requried")
	}
	if envs["SUBJECT"] != "" {
		b.Subject = envs["SUBJECT"]
	} else {
		return nil, fmt.Errorf("environment variable SUBJECT is requried")
	}
	if envs["TEXT"] != "" {
		b.Text = envs["TEXT"]
	} else {
		return nil, fmt.Errorf("environment variable TEXT is requried")
	}
	return b, nil
}

func (b *Builder) run() error {
	if err := b.SendEmail(); err != nil {
		return err
	}
	return nil
}

func (b *Builder) SendEmail() error {
	var toUsers = strings.Split(b.ToUsers, "|")
	m := gomail.NewMessage()
	//设置发件人
	m.SetAddressHeader("From", b.FromUser, "石头")
	//设置收件人
	m.SetHeader("To", toUsers...)
	//设置主题
	m.SetHeader("Subject", b.Subject)
	//设置正文
	//m.SetBody("text", "hello world!")
	m.SetBody("text/html", b.Text)
	//设置发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewPlainDialer("smtp.qq.com", 465, b.FromUser, b.Secret)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
	return nil
}

type CMD struct {
	Command []string
	WorkDir string
}

func (c CMD) Run() (string, error) {
	fmt.Println("Run CMD: ", strings.Join(c.Command, " "))

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

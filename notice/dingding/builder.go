package main

import (
	"fmt"
	"os/exec"
	"strings"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
)

const baseSpace = "/root/src"

type Builder struct {
	Webhook string
	AtMobiles []string
	IsAtAll bool
	Message string

	payload interface{}
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if b.Webhook = envs["WEBHOOK"]; b.Webhook == "" {
		return nil, fmt.Errorf("envionment variable WEBHOOK is required")
	}

	if envs["MESSAGE"] == ""  && envs["_WORKFLOW_TASK_DETAIL"] == ""{
		return nil, fmt.Errorf("envionment variable MESSAGE is required")
	}

	if envs["MESSAGE"] != "" {
		b.Message = envs["MESSAGE"]
		b.AtMobiles = strings.Split(envs["AT_MOBILES"], ",")

		if strings.ToLower(envs["IS_AT_ALL"]) == "true" {
			b.IsAtAll = true
		}

		text := Text {
			Content: b.Message,
		}
		at := At{
			AtMobiles: b.AtMobiles,
			IsAtAll: b.IsAtAll,
		}
		info := TextWebhook{
			Msgtype:  "text",
			Text:   text,
			At: at,
		}

		b.payload = info
		return b, nil
	}

	task := &FlowTask{}
	err := json.Unmarshal([]byte(envs["_WORKFLOW_TASK_DETAIL"]), task)
	if err != nil {
		return nil, err
	}
	var totalTime string
	if task.End != nil && task.Start != nil {
		totalTime = fmt.Sprintf(" 总耗时: %d 秒", (int64)(task.End.Sub(*task.Start).Seconds()))
	}

	link := Link{
		Title: fmt.Sprintf("工作流%s通知", task.Name),
		PicURL: "https://main.qcloudimg.com/raw/d8ff94e74414bdfb4474fe091608d53f.svg",
		MessageURL: task.DetailURL,
		Text: fmt.Sprintf("状态: %s\n%s", task.Status, totalTime),
	}

	b.payload = LinkWebhook{
		Msgtype:  "link",
		Link: link,
	}

	return b, nil

}

func (b *Builder) run() error {
	if err := b.callWebhook(); err != nil {
		return err
	}

	return nil
}

func (b *Builder) callWebhook() error {
	payload, _ := json.Marshal(b.payload)
	fmt.Printf("sending webhook info: %s\n", string(payload))
	body := bytes.NewBuffer(payload)
	res, err := http.Post(b.Webhook, "application/json;charset=utf-8", body)
	if err != nil {
		return err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	var resultJSON interface{}
	if err = json.Unmarshal(result, &resultJSON); err != nil {
	  return err
	}

	fmt.Println(resultJSON)
	fmt.Println("Send webhook succeded.")
	return nil
}

type CMD struct {
	Command []string // cmd with args
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

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
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if b.Webhook = envs["WEBHOOK"]; b.Webhook == "" {
		return nil, fmt.Errorf("envionment variable WEBHOOK is required")
	}

	b.AtMobiles = strings.Split(envs["AT_MOBILES"], ",")

	if strings.ToLower(envs["IS_AT_ALL"]) == "true" {
		b.IsAtAll = true
	}
	return b, nil
}

func (b *Builder) run() error {
	if err := b.callWebhook(); err != nil {
		return err
	}

	return nil
}

type TextWebhook struct {
	Msgtype string `json:"msgtype"`
	Text Text `json:"text"`
	At At `json:"at"`
}

type Text struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll bool `json:"isAtAll"`
}

func (b *Builder) callWebhook() error {
	text := Text {
		Content: "工作流执行结束",
	}
	at := At{
		AtMobiles: b.AtMobiles,
		IsAtAll: b.IsAtAll,
	}
	hookInfo := TextWebhook{
		Msgtype:  "text",
		Text:   text,
		At: at,
	}

	js, _ := json.Marshal(hookInfo)
	fmt.Printf("sending webhook info: %s\n", string(js))
	body := bytes.NewBuffer([]byte(js))
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
	// log.WithFields(log.Fields{"action": "service#UpdateGitHook", "event": "receiveGitToken"}).Debug(string(result))


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

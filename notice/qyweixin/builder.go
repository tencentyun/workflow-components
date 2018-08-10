package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

type Builder struct {
	CorpId    string
	AppSecret string
	Users     string
	Partys    string
	Tags      string
	AgentId   string
}

type TokenInfo struct {
	CorpId     string `json:"corp_id"`
	CorpSecret string `json:"corp_secret"`
}

type WebhookInfo struct {
	ToUser  string `json:"touser"`
	ToParty string `json:"toparty"`
	ToTag   string `json:"totag"`
	MsgType string `json:"msgtype"`
	AgentId int64  `json:"agentid"`
	Text    Text   `json:"text"`
	Safe    int64  `json:"safe"`
}

type Text struct {
	Content string `json:"content"`
}

type BackResponce struct {
	Errcode     int64  `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
type SendMsgReponce struct {
	Errcode      int64  `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if envs["CORPID"] != "" {
		b.CorpId = envs["CORPID"]
	} else {
		return nil, fmt.Errorf("environment variable CORPID is reuquired")
	}
	if envs["APPSECRET"] != "" {
		b.AppSecret = envs["APPSECRET"]
	} else {
		return nil, fmt.Errorf("environment variable APPSECRET is reuquired")
	}
	if envs["AGENTID"] != "" {
		b.AgentId = envs["AGENTID"]
	} else {
		return nil, fmt.Errorf("environment variable AGENTID is reuquired")
	}

	if envs["PARTYS"] != "" {
		b.Partys = envs["PARTYS"]
	}
	if envs["USERS"] != "" {
		b.Users = envs["USERS"]
	}
	if envs["TAGS"] != "" {
		b.Tags = envs["TAGS"]
	}

	if envs["USERS"] == "" && envs["PARTYS"] == "" && envs["TAGS"] == "" {
		return nil, fmt.Errorf("environment variable CORPID is reuquired")
	}

	return b, nil
}

func (b *Builder) run() error {
	token, err := b.GetToken()
	if err != nil {
		return err
	}
	err = b.SendMsg(token)
	if err != nil {
		return err
	}
	return nil
}

func (b *Builder) GetToken() (string, error) {
	var tokenInfo = TokenInfo{
		CorpId:     b.CorpId,
		CorpSecret: b.AppSecret,
	}
	//url = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", tokenInfo.CorpId, tokenInfo.CorpSecret)

	req, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	result, err := ioutil.ReadAll(req.Body)

	var resultJSON = &BackResponce{}
	if err = json.Unmarshal(result, resultJSON); err != nil {
		return "", err
	}

	fmt.Println(resultJSON.AccessToken)

	return resultJSON.AccessToken, nil
}

func (b *Builder) SendMsg(token string) error {
	text := Text{
		Content: "workflow haved finished",
	}

	agent, err := strconv.ParseInt(b.AgentId, 10, 64)
	if err != nil {
		return err
	}
	webhookInfo := WebhookInfo{
		ToUser:  b.Users,
		ToParty: b.Partys,
		ToTag:   b.Tags,
		MsgType: "text",
		AgentId: agent,
		Text:    text,
		Safe:    0,
	}

	js, _ := json.Marshal(webhookInfo)
	body := bytes.NewBuffer([]byte(js))
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)

	req, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return err
	}

	defer req.Body.Close()
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var resultJSON = &SendMsgReponce{}
	if err := json.Unmarshal(result, resultJSON); err != nil {
		return err
	}
	fmt.Println("send message success")
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

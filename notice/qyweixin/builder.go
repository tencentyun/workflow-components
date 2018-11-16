package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
)

type Builder struct {
	CorpId         string
	AppSecret      string
	Users          string
	Partys         string
	Tags           string
	AgentId        string
	MsgType        string
	Message        string
	PauseFlag      bool
	PauseResumeURL string
	PauseStopURL   string
	Payload        TextCard
}

func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if envs["CORP_ID"] == "" {
		return nil, fmt.Errorf("environment variable CORP_ID is reuquired")
	} else {
		b.CorpId = envs["CORP_ID"]
	}
	if envs["APP_SECRET"] == "" {
		return nil, fmt.Errorf("environment variable APP_SECRET is reuquired")
	} else {
		b.AppSecret = envs["APP_SECRET"]
	}
	if envs["AGENT_ID"] == "" {
		return nil, fmt.Errorf("environment variable AGENT_ID is reuquired")
	} else {
		b.AgentId = envs["AGENT_ID"]
	}

	if envs["USERS"] == "" && envs["PARTYS"] == "" && envs["TAGS"] == "" {
		return nil, fmt.Errorf("USERS OR PARTYS OR TAGS you must give one")
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

	b.PauseFlag = strings.ToLower(envs["_WORKFLOW_FLAG_PAUSE_NOTICE"]) == "true"

	if b.PauseFlag {
		b.PauseResumeURL = envs["_WORKFLOW_FLOW_PAUSE_HOOK_RESUME_API"]
		b.PauseStopURL = envs["_WORKFLOW_FLOW_PAUSE_HOOK_STOP_API"]
	}

	if envs["MESSAGE"] != "" {
		b.MsgType = "text"
		b.Message = envs["MESSAGE"]

		if b.PauseFlag {
			b.Message += fmt.Sprintf("\n\n当前工作流处于暂停状态\n继续执行链接:%s\n终止执行链接:%s", b.PauseResumeURL, b.PauseStopURL)
		}
		return b, nil
	}

	task := &FlowTask{}
	err := json.Unmarshal([]byte(envs["_WORKFLOW_TASK_DETAIL"]), task)
	if err != nil {
		return nil, err
	}

	taskNew := getFlowTaskNew(task, b.PauseFlag, b.PauseResumeURL, b.PauseStopURL)
	var totalTime string
	if taskNew.End != nil && taskNew.Start != nil {
		totalTime = fmt.Sprintf("总耗时: %d 秒", (int64)(taskNew.End.Sub(*taskNew.Start).Seconds()))
	}

	description := fmt.Sprintf(`<div class=\"normal\">状态: %s<br>%s</div>`, taskNew.Status, totalTime)

	if b.PauseFlag {
		description += fmt.Sprintf(`<div class=\"gray\">当前工作流处于暂停状态</div><div class=\"gray\">继续执行链接:%s<br>终止执行链接:%s</div>`, b.PauseResumeURL, b.PauseStopURL)
	}

	textCard := TextCard{
		Title:       fmt.Sprintf("%s/%s/%s", taskNew.Namespace, taskNew.Repo, taskNew.Name),
		MessageURL:  taskNew.DetailURL,
		Description: description,
		BtnTxt:      "详情",
	}
	b.MsgType = "textcard"
	b.Payload = textCard

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

func getFlowTaskNew(task *FlowTask, flag bool, resume, stop string) *FlowTaskNew {
	return &FlowTaskNew{
		Namespace:      task.Namespace,
		Repo:           task.Repo,
		Name:           task.Name,
		Status:         task.Status,
		Start:          task.Start,
		End:            task.End,
		DetailURL:      task.DetailURL,
		PauseFlag:      flag,
		PauseResumeURL: resume,
		PauseStopURL:   stop,
	}
}

func (b *Builder) GetToken() (string, error) {
	var tokenInfo = TokenInfo{
		CorpId:     b.CorpId,
		CorpSecret: b.AppSecret,
	}

	var param = make(url.Values)
	param.Add("corpid", tokenInfo.CorpId)
	param.Add("corpsecret", tokenInfo.CorpSecret)

	var paramStr = param.Encode()

	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?" + paramStr

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

	//fmt.Println(resultJSON.AccessToken)

	return resultJSON.AccessToken, nil
}

func (b *Builder) SendMsg(token string) error {
	agent, err := strconv.ParseInt(b.AgentId, 10, 64)
	if err != nil {
		return err
	}
	var webhookInfo interface{}
	if b.MsgType == "text" {
		text := Text{
			Content: b.Message,
		}
		webhookInfo = TextWebhook{
			ToUser:  b.Users,
			ToParty: b.Partys,
			ToTag:   b.Tags,
			MsgType: "text",
			AgentId: agent,
			Text:    text,
			Safe:    0,
		}
	} else if b.MsgType == "textcard" {
		webhookInfo = LinkWebhook{
			ToUser:   b.Users,
			ToParty:  b.Partys,
			ToTag:    b.Tags,
			MsgType:  "textcard",
			AgentId:  agent,
			TextCard: b.Payload,
		}
	}

	js, _ := json.Marshal(webhookInfo)
	body := bytes.NewBuffer([]byte(js))

	var param = make(url.Values)
	param.Add("access_token", token)
	var paramStr = param.Encode()

	var url = "https://qyapi.weixin.qq.com/cgi-bin/message/send?" + paramStr
	//fmt.Println(url)

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
	if resultJSON.Errcode != 0 {
		return fmt.Errorf("bad result: %s", result)
	}
	fmt.Printf("qyweixin return msg:\nerrcode: %d\nerrmsg: %s\ninvaliduser: %s\ninvalidparty: %s\ninvalidtag: %s\n", resultJSON.Errcode, resultJSON.Errmsg, resultJSON.InvalidUser, resultJSON.InvalidParty, resultJSON.InvalidTag)
	fmt.Println("send msg success!")

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

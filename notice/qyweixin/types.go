package main

import "time"

type FlowTask struct {
	Namespace string     `json:"namespace"`
	Repo      string     `json:"repo"`
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	Start     *time.Time `json:"start,omitempty"`
	End       *time.Time `json:"end,omitempty"`
	DetailURL string     `json:"detail_url"`
}

type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	MessageURL  string `json:"url"`
	BtnTxt      string `json:"btntxt"`
}

type TokenInfo struct {
	CorpId     string `json:"corp_id"`
	CorpSecret string `json:"corp_secret"`
}
type Text struct {
	Content string `json:"content"`
}
type TextWebhook struct {
	ToUser  string `json:"touser"`
	ToParty string `json:"toparty"`
	ToTag   string `json:"totag"`
	MsgType string `json:"msgtype"`
	AgentId int64  `json:"agentid"`
	Text    Text   `json:"text"`
	Safe    int64  `json:"safe"`
}
type LinkWebhook struct {
	ToUser   string   `json:"touser"`
	ToParty  string   `json:"toparty"`
	ToTag    string   `json:"totag"`
	MsgType  string   `json:"msgtype"`
	AgentId  int64    `json:"agentid"`
	TextCard TextCard `json:"textcard"`
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

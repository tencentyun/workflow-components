package main

import "time"

type Stage struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
	Jobs   []Job  `json:"jobs"`
}

type Job struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type FlowTask struct {
	Namespace string     `json:"namespace"`
	Repo      string     `json:"repo"`
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	Start     *time.Time `json:"start,omitempty"`
	End       *time.Time `json:"end,omitempty"`
	DetailURL string     `json:"detail_url"`
	Stages    []Stage    `json:"stages"`
}

type LinkWebhook struct {
	Msgtype string `json:"msgtype"`
	Link    Link   `json:"link"`
}

type Link struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	PicURL     string `json:"picUrl"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
}

type TextWebhook struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
	At      At     `json:"at"`
}

type Text struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type MarkdownWebHook struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

package main

import "time"

type FlowTask struct {
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	Start     *time.Time `json:"start,omitempty"`
	End       *time.Time `json:"end,omitempty"`
	DetailURL string     `json:"detail_url"`
}

type LinkWebhook struct {
	Msgtype string `json:"msgtype"`
	Link Link `json:"link"`
}

type Link struct {
	ID string `json:"id"`
	Title string `json:"title"`
	PicURL string `json:"picUrl"`
	Text string `json:"text"`
	MessageURL string `json:"messageUrl"`
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


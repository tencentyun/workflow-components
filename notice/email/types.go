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

type EmailContent struct {
	Title string
	Text  string
	Url   string
}

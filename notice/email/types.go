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
type FlowTaskNew struct {
	Namespace      string     `json:"namespace"`
	Repo           string     `json:"repo"`
	Name           string     `json:"name"`
	Status         string     `json:"status"`
	Start          *time.Time `json:"start,omitempty"`
	End            *time.Time `json:"end,omitempty"`
	DetailURL      string     `json:"detail_url"`
	Stages         []Stage    `json:"stages"`
	PauseFlag      bool       `json:"pause_flag"`
	PauseResumeURL string     `json:"pause_resume_url"`
	PauseStopURL   string     `json:"pause_stop_url"`
}

type EmailContent struct {
	Title string
	Text  string
	Url   string
}

package main

type StartBy struct {
	StartType string `json:"start_type"`
	Time      string `json:"time"`
}
type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Resource struct {
	MaxCnt int    `json:"max_cnt"`
	MinCnt int    `json:"min_cnt"`
	Type   string `json:"type"`
	Group  string `json:"group"`
}
type CreateTest struct {
	TestTypeID int        `json:"testtype_id"`
	Name       string     `json:"name"`
	StartBy    StartBy    `json:"start_by"`
	ProjectID  int        `json:"project_id"`
	Properties []Property `json:"properties"`
	Resources  []Resource `json:"resources"`
}

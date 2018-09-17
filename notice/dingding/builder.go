package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

const STAGE_TYPE_END = "end"

const baseSpace = "/root/src"

type Builder struct {
	Webhook   string
	AtMobiles []string
	IsAtAll   bool
	Message   string

	payload interface{}
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if b.Webhook = envs["WEBHOOK"]; b.Webhook == "" {
		return nil, fmt.Errorf("envionment variable WEBHOOK is required")
	}

	if envs["MESSAGE"] == "" && envs["_WORKFLOW_TASK_DETAIL"] == "" {
		return nil, fmt.Errorf("envionment variable MESSAGE is required")
	}

	if strings.ToLower(envs["IS_AT_ALL"]) == "true" {
		b.IsAtAll = true
	}

	b.AtMobiles = strings.Split(envs["AT_MOBILES"], ",")
	at := At{
		AtMobiles: b.AtMobiles,
		IsAtAll:   b.IsAtAll,
	}

	if envs["MESSAGE"] != "" {
		b.Message = envs["MESSAGE"]

		text := Text{
			Content: b.Message,
		}

		info := TextWebhook{
			Msgtype: "text",
			Text:    text,
			At:      at,
		}

		b.payload = info
		return b, nil
	}

	task := &FlowTask{}
	err := json.Unmarshal([]byte(envs["_WORKFLOW_TASK_DETAIL"]), task)

	if err != nil {
		return nil, err
	}

	//use template
	t := template.New("taskflow")
	t.Funcs(template.FuncMap{"myFunc": showStage, "totalTime": timeConsuming})
	t, _ = t.Parse(`### {{.Name}}{{"\n"}}状态: {{ .Status}}{{"\n"}}{{totalTime .Start .End}}{{"\n"}}
		{{range .Stages}}
    		{{myFunc .}}
		{{end}}
		{{"\n"}}[点击链接查看详情]({{.DetailURL}})
		`)

	buf := new(bytes.Buffer)
	t.Execute(buf, task)
	fmt.Println(buf.String())

	md := Markdown{
		Title: "工作流通知",
		Text:  buf.String(),
	}

	b.payload = MarkdownWebHook{
		Msgtype:  "markdown",
		Markdown: md,
		At:       at,
	}

	return b, nil

}
func showStage(stage Stage) string {
	var mdText = "\n"
	nm := stage.Name
	status := stage.Status
	stageType := stage.Type
	jobs := stage.Jobs
	if stageType != STAGE_TYPE_END {
		mdText += fmt.Sprintf("- **%s** : **%s** \n", nm, status)

		for _, job := range jobs {
			name := job.Name
			status := job.Status
			mdText += fmt.Sprintf("> - **%s** : **%s** \n", name, status)
		}
	}
	return mdText
}

func timeConsuming(start *time.Time, end *time.Time) string {
	var totalTime string
	if start != nil && end != nil {
		totalTime = fmt.Sprintf("总耗时: %d 秒", (int64)(end.Sub(*start).Seconds()))
	}
	return totalTime
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
	fmt.Println("Send webhook succeed.")
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

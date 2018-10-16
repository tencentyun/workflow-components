package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/imroc/req"
	"golang.org/x/net/context"
)

const (
	APP_ID  = "3321337994"
	QTA_URL = "http://qta.tencentyun.com"
	RUN_URL = "http://qta.tencentyun.com/api/v1/task/task/"
)

type Builder struct {
	PlanID    int
	TaskID    int
	Name      string
	Puin      string
	endStatus string
	reportID  string
	payload   interface{}
}

func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if envs["PLAN_ID"] == "" && envs["_WORKFLOW_TASK_PLAN_ID"] == "" {
		return nil, fmt.Errorf("environment variable PLAN_ID is required")
	}

	if envs["PLAN_ID"] != "" {
		b.PlanID, _ = strconv.Atoi(envs["PLAN_ID"])
	} else {
		b.PlanID, _ = strconv.Atoi(envs["_WORKFLOW_TASK_PLAN_ID"])
	}
	b.Puin = envs["_WORKFLOW_FLOW_UIN"]
	fmt.Printf("uin:%s appid:%s\n", b.Puin, APP_ID)
	return b, nil
}

type LoopResult struct {
	Data string
	Err  error
}

func (b *Builder) run() error {
	if err := b.startTqa(); err != nil {
		return err
	}

	if err := b.doLoop(); err != nil {
		return err
	}

	return nil
}

func (b *Builder) startTqa() error {
	r := req.New()
	r.SetTimeout(10 * time.Second)
	body := req.BodyJSON(map[string]interface{}{"plan_id": b.PlanID})
	header := req.Header{
		"Authorization": fmt.Sprintf("TencentHub %s skey %s", b.Puin, APP_ID),
	}
	resp, err := r.Post(RUN_URL, header, body)
	if err != nil {
		return fmt.Errorf("read http response failed:%v", err)
	}

	fmt.Printf("body:%s\n", resp.String())

	var plan struct {
		ID     int    `json:"id"`
		PlanID int    `json:"plan_id"`
		Name   string `json:"name"`
	}

	err = resp.ToJSON(&plan)
	if err != nil {
		return fmt.Errorf("unmarshal resp failed:%v", err)
	}

	b.TaskID = plan.ID
	return nil
}

//超时时间设置为30分钟，轮询间隔设置为2分钟
func (b *Builder) doLoop() error {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Minute)
	status, err := b.loopDoRequest(ctx)
	if err != nil {
		return err
	}
	if status == "timeout" {
		fmt.Printf("The preset 30-minute query time has timed out. For the result of task running, please log in to qta.")
	} else {
		if err := b.showQtaReport(); err != nil {
			return err
		}
	}
	return nil
}

func (b *Builder) loopDoRequest(ctx context.Context) (string, error) {
	for {
		res := b.doReq()
		if res == "fail" {
			return "fail", fmt.Errorf("")
		} else if res == "finish" {
			return res, nil
		}
		select {
		case <-ctx.Done():
			return "timeout", nil
		default:
			time.Sleep(2 * time.Minute)
		}
	}
}

//curl -v -H "Content-Type: application/json" -H "Authorization: TencentHub 3321337994 skey 3321337994" http://qta.tencentyun.com/api/v1/task/task/5617
func (b *Builder) doReq() string {
	queryURL := fmt.Sprintf("http://qta.tencentyun.com/api/v1/task/task/%d/", b.TaskID)
	fmt.Println(queryURL)
	r := req.New()
	r.SetTimeout(10 * time.Second)
	header := req.Header{
		"Authorization": fmt.Sprintf("TencentHub %s skey %s", b.Puin, APP_ID),
	}
	resp, err := r.Get(queryURL, header)
	if err != nil {
		fmt.Printf("read http response failed:%v", err)
		return "fail"
	}
	if resp.Response().StatusCode != http.StatusOK {
		fmt.Printf("bad status code: %s, body: %s", resp.Response().Status, resp.String())
		return "fail"
	}

	fmt.Printf("resp from QTA:%s\n", resp.String())

	var queryRes struct {
		Status   string `json:"status"`
		ReportID string `json:"report_id"`
	}
	if err = resp.ToJSON(&queryRes); err != nil {
		fmt.Printf("unmarshal resp failed:%v", err)
		return "fail"
	}
	fmt.Printf("the result of task running: %s\n", queryRes.Status)
	if queryRes.Status == "init" || queryRes.Status == "waiting" || queryRes.Status == "running" {
		return "unfinish"
	}
	b.endStatus = queryRes.Status
	b.reportID = queryRes.ReportID
	return "finish"
}
func (b *Builder) showQtaReport() error {
	//当任务执行的最终状态不为done或者report_id为空时，直接返回，不需要打印报告
	if b.endStatus != "done" || b.reportID == "" {
		return nil
	}
	queryURL := fmt.Sprintf("http://qta.tencentyun.com/api/v1/report/%s/case/", b.reportID)
	fmt.Println(queryURL)
	r := req.New()
	r.SetTimeout(10 * time.Second)
	header := req.Header{
		"Authorization": fmt.Sprintf("TencentHub %s skey %s", b.Puin, APP_ID),
	}
	resp, err := r.Get(queryURL, header)
	if err != nil {
		return fmt.Errorf("do http req failed:%v", err)
	}
	if resp.Response().StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code: %s, body: %s", resp.Response().Status, resp.String())
	}

	type testCase struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Result   string `json:"result"`
		ReportID int    `json:"report_id"`
		Reason   string `json:"reason"`
	}

	var queryRes struct {
		Count   int        `json:"count"`
		Results []testCase `json:"results"`
	}

	if err = resp.ToJSON(&queryRes); err != nil {
		return fmt.Errorf("unmarshal resp failed:%v", err)
	}

	fmt.Printf("report from QTA:%+v\n", queryRes)

	fmt.Println("follows, urls report more test case detail:")
	for _, tc := range queryRes.Results {
		detail := fmt.Sprintf("%s/api/v1/report/%s/case/%d/info/html/", QTA_URL, b.reportID, tc.ID)
		fmt.Println(detail)
	}

	return nil
}

//curl -v -H "Content-Type:application/json" -H "Authorization: TencentHub 3105311399 skey 251005787" -X POST --data '{"name": "test04", "testtype_id": 3, "start_by": {"start_type": "manual"}, "project_id": 1, "properties": [{"name": "TEST_REPO_URL", "value": "https://drun-1256586152.cos.ap-guangzhou.myqcloud.com/qt4a-debug.tgz"}, {"name": "PRODUCT_PATH", "value": "https://drun-1256586152.cos.ap-guangzhou.myqcloud.com/app-debug-2.apk"}, {"name": "TESTCASENAME", "value": "demotest"}], "resources": [{"max_cnt": 1, "min_cnt": 1, "type": "node", "group": "NodePool"}, {"max_cnt": 1, "min_cnt": 1, "type": "android", "group": "QT4A_Cloudroid"}]}'  --cookie "app_id=251005787;p_uin=o3105311399;skey=EtUUy2cgmJV1*ylTrjziV-muNxA188GmMc2JrNT*pGs_" http://qta.tencentyun.com/api/v1/task/plan/
//curl -v -H "Content-Type:application/json" -H "Authorization: TencentHub 3321337994 skey 3321337994" -X POST --data '{"plan_id": 28}' http://qta.tencentyun.com/api/v1/task/task/

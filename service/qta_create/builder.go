package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	APP_ID = "3321337994"
	URL    = "http://qta.tencentyun.com/api/v1/task/plan/"
)

type Builder struct {
	TestTypeID int
	Name       string
	Puin       string
	StartType  string
	StartTime  string
	ProjectID  int
	payload    interface{}
}

func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	// if envs["TEST_TYPE_ID"] == "" {
	// 	return nil, fmt.Errorf("environment variable TEST_TYPE_ID is required")
	// }
	if envs["NAME"] == "" {
		return nil, fmt.Errorf("environment variable NAME is required")
	}
	if envs["PRODUCT_PATH"] == "" {
		return nil, fmt.Errorf("environment variable PRODUCT_PATH is required")
	}
	if envs["TEST_REPO_URL"] == "" {
		return nil, fmt.Errorf("environment variable TEST_REPO_URL is required")
	}
	if envs["TESTCASENAME"] == "" {
		return nil, fmt.Errorf("environment variable TESTCASENAME is required")
	}
	// if envs["START_TYPE"] == "" {
	// 	return nil, fmt.Errorf("environment variable START_TYPE is required")
	// }

	// b.TestTypeID, _ = strconv.Atoi(envs["TEST_TYPE_ID"])
	// b.ProjectID, _ = strconv.Atoi(envs["PROJECT_ID"])
	b.Name = envs["NAME"]
	b.Puin = envs["_WORKFLOW_FLOW_UIN"]
	fmt.Printf("uin:%s\n", b.Puin)
	start := StartBy{
		StartType: "manual",
	}
	//属性列表
	properties := make([]Property, 0)

	property01 := Property{
		Name:  "PRODUCT_PATH",
		Value: envs["PRODUCT_PATH"],
	}
	properties = append(properties, property01)

	property02 := Property{
		Name:  "TEST_REPO_URL",
		Value: envs["TEST_REPO_URL"],
	}
	properties = append(properties, property02)

	property03 := Property{
		Name:  "TESTCASENAME",
		Value: envs["TESTCASENAME"],
	}
	properties = append(properties, property03)

	//资源列表
	resources := make([]Resource, 0)
	resource01 := Resource{
		MaxCnt: 1,
		MinCnt: 1,
		Type:   "node",
		Group:  "NodePool",
	}
	resources = append(resources, resource01)

	resource02 := Resource{
		MaxCnt: 1,
		MinCnt: 1,
		Type:   "android",
		Group:  "QT4A_Cloudroid",
	}
	resources = append(resources, resource02)

	createTest := CreateTest{
		TestTypeID: 3,
		Name:       b.Name,
		StartBy:    start,
		ProjectID:  1,
		Properties: properties,
		Resources:  resources,
	}

	b.payload = createTest
	return b, nil
}

func (b *Builder) run() error {
	if err := b.callCreateTestPlan(); err != nil {
		return err
	}
	return nil
}
func (b *Builder) callCreateTestPlan() error {
	payload, _ := json.Marshal(b.payload)
	fmt.Printf("sending qta create info: %s\n", string(payload))
	query := bytes.NewBuffer(payload)
	httpRequest, err := http.NewRequest("POST", URL, query)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("TencentHub %s skey %s", b.Puin, APP_ID))

	http.DefaultClient.Timeout = time.Duration(time.Duration(10) * time.Second)
	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return fmt.Errorf("do http failed:%v", err)
	}

	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return fmt.Errorf("read http response failed:%v", err)
	}
	fmt.Printf("resp: %v\n", string(body))
	type plan struct {
		PlanId int    `json:"id"`
		Name   string `json:"name"`
	}

	var resp plan
	if err = json.Unmarshal(body, &resp); err != nil {
		return fmt.Errorf("unmarshal resp failed:%v", err)
	}
	fmt.Printf("[JOB_OUT] _WORKFLOW_TASK_PLAN_ID = %s\n", strconv.Itoa(resp.PlanId))

	return nil
}

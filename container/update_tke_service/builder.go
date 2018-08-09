package main

import (
	"fmt"
	"strings"
	"time"
	"net/http"
	"io/ioutil"
	"sort"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"encoding/json"
	"errors"
	"context"
)

// Builder is
type Builder struct {
	SecretID string
	SecretKey string
	ClusterID string
	ServiceName string
	Namespace string
	Region string
	Image string
	Containers map[string]string
}

const TIMEOUT = 5

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}
	if envs["TENCENTCLOUD_SECRET_KEY"] == "" {
		return nil, fmt.Errorf("envionment variable TENCENTCLOUD_SECRET_KEY is required")
	}
	b.SecretKey = envs["TENCENTCLOUD_SECRET_KEY"]

	if envs["TENCENTCLOUD_SECRET_ID"] == "" {
		return nil, fmt.Errorf("envionment variable TENCENTCLOUD_SECRET_ID is required")
	}
	b.SecretID = envs["TENCENTCLOUD_SECRET_ID"]

	if envs["CLUSTER_ID"] == "" {
		return nil, fmt.Errorf("envionment variable CLUSTER_ID is required")
	}
	b.ClusterID = envs["CLUSTER_ID"]

	if envs["SERVICE_NAME"] == "" {
		return nil, fmt.Errorf("envionment variable SERVICE_NAME is required")
	}
	b.ServiceName = envs["SERVICE_NAME"]

	if envs["REGION"] == "" {
		return nil, fmt.Errorf("envionment variable REGION is required")
	}
	b.Region = envs["REGION"]

	if envs["IMAGE"] == "" && envs["CONTAINERS"] == "" {
		return nil, fmt.Errorf("envionment variable IMAGE or CONTAINERS is required")
	}

	if envs["CONTAINERS"] != "" {
		var containers map[string]string
		if err := json.Unmarshal([]byte(envs["CONTAINERS"]), &containers); err != nil {
			return nil, fmt.Errorf("envionment variable CONTAINERS must be valid json string")
		}
		b.Containers = containers
	} else {
		b.Image = envs["IMAGE"]
		if strings.Index(b.Image, ":") == -1 { // add default tag
			b.Image = fmt.Sprintf("%s:latest", b.Image)
		}
	}

	if envs["NAMESPACE"] == "" {
		envs["NAMESPACE"] = "default"
	}
	b.Namespace = envs["NAMESPACE"]


	return b, nil
}

func (b *Builder) run() (err error) {
	err = b.updateService()
	if err != nil {
		return err
	}

	// time.Sleep(10*time.Second)
	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Duration(TIMEOUT) * time.Minute)
	// ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Duration(TIMEOUT) * time.Second)
	defer cancelFunc()

	c := make(chan struct{})
	go func() {
		defer func() {
			close(c)
		}()
	    for {
			ok, e := b.checkService()
			if e != nil {
				err = e
				return
			}
			if ok {
				return
			}
			time.Sleep(10*time.Second)
		}
	}()

	select {
	case <-ctx.Done():
		err = errors.New("获取服务状态超时, 请手动检查服务是否更新成功")
		return
	case <-c:
		return
	}
}

func (b *Builder) updateService() error {
	params := map [string]string{
		"clusterId": b.ClusterID,
		"serviceName": b.ServiceName,
		"image": b.Image,
		"namespace": b.Namespace,
		"Action": "ModifyClusterServiceImage",
	}

	if b.Image != "" {
		params["image"] = b.Image
	} else {
		var index int64
		for k, v := range b.Containers {
			params[fmt.Sprintf("containers.%d.containerName", index)] = k
			params[fmt.Sprintf("containers.%d.image", index)] = v
			index++
		}
	}

	body, err := b.callAPI(params)
	if err != nil {
		return err
	}

	resultJSON := &struct {
		Code int64
		CodeDesc string
		Message string
	}{}
	if err = json.Unmarshal(body, resultJSON); err != nil {
		return err
	}

	fmt.Println(resultJSON)

	if resultJSON.Code == 0 && resultJSON.CodeDesc == "Success" {
		return nil
	}
	return errors.New(resultJSON.Message)
}

func (b *Builder) checkService() (bool, error) {
	params := map [string]string{
		"clusterId": b.ClusterID,
		"serviceName": b.ServiceName,
		"namespace": b.Namespace,
		"Action": "DescribeClusterServiceInfo",
	}

	body, err := b.callAPI(params)
	if err != nil {
		return false, err
	}

	resultJSON := &struct {
		Code int64
		CodeDesc string
		Message string
		Data struct {
			Service struct{
				ServiceName string
				Status string
				ReasonMap map[string]int64
				Containers []struct {
					ContainerName string
					Image string
				}
			}
		}
	}{}
	if err = json.Unmarshal(body, resultJSON); err != nil {
		return false, err
	}

	fmt.Println(resultJSON)

	if resultJSON.Code == 0 && resultJSON.CodeDesc == "Success" {
		fmt.Println("服务更新状态: ", resultJSON.Data.Service.Status)
		if resultJSON.Data.Service.Status == "Normal" {
			return true, nil
		}
		return false, nil
	}
	return false, errors.New(resultJSON.Message)
}


func (b *Builder) callAPI(actionParams map[string]string) ([]byte, error) {
	t := time.Now().Unix()
	rand.Seed(t)
	randNum := rand.Intn(1000)

	params := map [string]string{
		"SecretId": b.SecretID,
		"Region": b.Region,
		"Timestamp": fmt.Sprintf("%d",t),
		"Nonce": fmt.Sprintf("%d",randNum),
		"SignatureMethod": "HmacSHA256",
	}

	for k, v := range actionParams {
		params[k] = v
	}

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	req, err := http.NewRequest("GET", "https://ccs.api.qcloud.com/v2/index.php", nil)
	if err != nil {
		return nil, err
	}

	var args []string
	q := req.URL.Query()
	for _, k := range keys {
		q.Add(k, params[k])
		args = append(args, fmt.Sprintf("%s=%s", k, params[k]))
	}

	url := "GETccs.api.qcloud.com/v2/index.php?" + strings.Join(args, "&")

	signStr := sign(url, b.SecretKey)
	q.Add("Signature", signStr)

	req.URL.RawQuery = q.Encode()
	//fmt.Println("Call Api: ")
	//fmt.Println(req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// fmt.Println(resp.Status)
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func sign(s, secretKey string) string {
	hashed := hmac.New(sha256.New, []byte(secretKey))
	hashed.Write([]byte(s))

	return base64.StdEncoding.EncodeToString(hashed.Sum(nil))
}

package main

import (
	"fmt"
	// "time"
	"os"
)

var envList = []string{
	"TENCENTCLOUD_SECRET_ID",
	"TENCENTCLOUD_SECRET_KEY",
	"REGION",
	"CLUSTER_ID",
	"SERVICE_NAME",
	"CONTAINERS",
	"IMAGE",
	"NAMESPACE",
}

func main() {
	envs := make(map[string]string)
	for _, envName := range envList {
		envs[envName] = os.Getenv(envName)
	}

	builder, err := NewBuilder(envs)
	if err != nil {
		fmt.Println("UPDATE FAILED: ", err)
		os.Exit(1)
	}

	if err := builder.run(); err != nil {
		fmt.Println("UPDATE FAILED: ", err)
		os.Exit(1)
	}
	fmt.Println("UPDATE SUCCEED")
}

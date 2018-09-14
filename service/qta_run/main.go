package main

import (
	"fmt"
	"os"
)

var envList = []string{
	"PLAN_ID",
	"_WORKFLOW_TASK_PLAN_ID",
	"_WORKFLOW_FLOW_UIN",
}

func main() {
	envs := make(map[string]string)
	for _, nm := range envList {
		envs[nm] = os.Getenv(nm)
	}

	builder, err := NewBuilder(envs)
	if err != nil {
		fmt.Printf("build fail:%v\n", err)
		os.Exit(1)
	}

	if err := builder.run(); err != nil {
		fmt.Printf("build fail:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("build success")
}

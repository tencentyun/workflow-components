package main

import (
	"fmt"
	"os"
)

var envList = []string{
	"WEBHOOK",
	"AT_MOBILES",
	"IS_AT_ALL",
	"MESSAGE",
	"_WORKFLOW_TASK_DETAIL",
	"_WORKFLOW_FLAG_PAUSE_NOTICE",
	"_WORKFLOW_FLOW_PAUSE_HOOK_RESUME_API",
	"_WORKFLOW_FLOW_PAUSE_HOOK_STOP_API",
}

func main() {
	envs := make(map[string]string)
	for _, envName := range envList {
		envs[envName] = os.Getenv(envName)
	}

	builder, err := NewBuilder(envs)
	if err != nil {
		fmt.Println("BUILDER FAILED: ", err)
		os.Exit(1)
	}

	if err := builder.run(); err != nil {
		fmt.Println("BUILD FAILED", err)
		os.Exit(1)
	} else {
		fmt.Println("BUILD SUCCEED")
	}
}

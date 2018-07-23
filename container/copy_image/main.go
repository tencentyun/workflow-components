package main

import (
	"fmt"
	// "time"
	"os"
	"strings"
)

var envList = []string{
	"IMAGE_TAG",

	"HUB_USER",
	"HUB_TOKEN",
	"_WORKFLOW_HUB_USER",
	"_WORKFLOW_HUB_TOKEN",

	"TO_IMAGE",
	"TO_HUB_USER",
	"TO_HUB_TOKEN",
}

func main() {
	envs := make(map[string]string)
	for _, envName := range envList {
		envs[envName] = strings.TrimSpace(os.Getenv(envName))
	}

	builder, err := NewBuilder(envs)
	if err != nil {
		fmt.Println("BUILDER FAILED: ", err)
		os.Exit(1)
	}

	if err := builder.run(); err != nil {
		fmt.Println("BUILD FAILED: ", err)
		os.Exit(1)
	} else {
		fmt.Println("BUILD SUCCEED")
	}
}

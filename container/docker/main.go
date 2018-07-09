package main

import (
	"fmt"
	// "time"
	"os"
)

var envList = []string{
	"GIT_CLONE_URL",
	"GIT_REF",
	"GIT_TYPE",

	"IMAGE",
	"IMAGE_TAG_FORMAT",
	"IMAGE_TAG",
	"EXTRA_IMAGE_TAG",

	"BUILD_WORKDIR",
	"DOCKERFILE_PATH",
	"BUILD_ARGS",

	"HUB",
	"HUB_USER",
	"HUB_TOKEN",

	// "GIT_URL",
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
		fmt.Println("BUILD FAILED: ", err)
		os.Exit(1)
	} else {
		fmt.Println("BUILD SUCCEED")
	}
}

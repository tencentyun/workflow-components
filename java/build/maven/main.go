package main

import (
	"fmt"
	"os"
)

var envList = []string{"GIT_CLONE_URL", "GIT_REF", "GOALS", "POM_PATH"}

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

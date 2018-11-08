package main

import (
	"fmt"
	"os"
)

var envList = []string{
	"CACHE_ID",
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
		fmt.Println("CLEAN FAILED: ", err)
		os.Exit(1)
	} else {
		fmt.Println("CLEAN SUCCEED")
	}
}

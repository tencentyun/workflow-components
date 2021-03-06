package main

import (
	"fmt"
	"os"
)

var envsList = []string{
	"GIT_CLONE_URL",
	"GIT_REF",
	"OUT_PUT_XML",
	"_WORKFLOW_GIT_CLONE_URL",
	"_WORKFLOW_GIT_REF",

	"_WORKFLOW_FLAG_CACHE",
}

func main() {
	envs := make(map[string]string)
	for _, nm := range envsList {
		envs[nm] = os.Getenv(nm)
	}

	builder, err := NewBuilder(envs)
	if err != nil {
		fmt.Println("build failed", err)
		os.Exit(1)
	}

	if err := builder.run(); err != nil {
		fmt.Println("build failed", err)
		os.Exit(1)
	}

	fmt.Println("build success")
}

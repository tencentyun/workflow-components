package main

import (
	"fmt"
	"os"
)

var envsList = []string{
	"GIT_CLONE_URL",
	"GIT_REF",
	"WORKFLOW_GIT_CLONE_URL",
	"WORKFLOW_GIT_REF",
	"GTEST_PACKAGE_FILE",
	"GTEST_PARAMS",
}

func main() {
	envs := make(map[string]string)

	for _, nm := range envsList {
		envs[nm] = os.Getenv(nm)
	}

	builder, err := NewBuilder(envs)
	if err != nil {
		fmt.Println("build failed: ", err)
		os.Exit(1)
	}

	if err := builder.run(); err != nil {
		fmt.Println("build fail: ", err)
		os.Exit(1)
	}
	fmt.Println("build seccess")
}

package main

import (
	"fmt"
	"os"
)

var envList = []string{
	"USERNAME",
	"PASSWORD",
	"CERTIFICATE",
	"SERVER",
	// "RESOURCE",

	"CMD",
}

func main() {
	envs := make(map[string]string)
	for _, envName := range envList {
		envs[envName] = os.Getenv(envName)
	}

	//for _, e := range os.Environ() {
	//	pair := strings.Split(e, "=")
	//	k, v := pair[0], pair[1]
	//	if strings.HasPrefix(k, "ARG_") {
	//
	//	}
	//}

	builder, err := NewBuilder(envs)
	if err != nil {
		fmt.Println("BUILDER FAILED: ", err)
		os.Exit(1)
	}

	if err := builder.run(); err != nil {
		fmt.Println("KUBECTL FAILED: ", err)
		os.Exit(1)
	} else {
		fmt.Println("KUBECTL  SUCCEED")
	}
}

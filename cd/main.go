package main

import (
	"fmt"
	"os"
	"strings"
)

var envList = []string{
	"USERNAME",
	"PASSWORD",
	"CERTIFICATE",
	"SERVER",
	"NAMESPACE",
	"ACTION",

	// deploy
	"DEPLOY_GROUP",
	"DEPLOY_TARGET",
	"DEPLOYMENT_NAME",
	"REPLICAS",
	"STRATEGY",
	"IMAGE",
	"SERVICES",

	// Scale
	"DEPLOYMENT_NAME",
	"DEPLOY_TARGET",
	"DEPLOY_GROUP",
	"SCALE_TO",
	"SCALE_UP",
	"SCALE_DOWN",
	"AUTO_DELETION",

	// Disable
	"DEPLOYMENT_NAME",
	"DEPLOY_GROUP",
	"DEPLOY_TARGET",

	// delete
	"DEPLOYMENT_NAME",
	"DEPLOY_GROUP",
	"DEPLOY_TARGET",

	// shrink
	"DEPLOY_GROUP",
	"SHRINK_TO",

	// Enable
	"DEPLOYMENT_NAME",
	"DEPLOY_TARGET",

	// Rollback
	"FROM_DEPLOYMENT_NAME",
	"FROM_DEPLOY_TARGET",
	"TO_DEPLOYMENT_NAME",
	"TO_DEPLOY_TARGET",
	"DEPLOY_GROUP",
	// "VERSION",
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

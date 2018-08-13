package main

import (
	"fmt"
	"os"
)

var envList = []string{
	"FROM_USER",
	"TO_USERS",
	"SECRET",
	"SUBJECT",
	"SMTP_SERVER",
	"SMTP_PORT",
	"TEXT",
}

func main() {
	envs := make(map[string]string)

	for _, name := range envList {
		envs[name] = os.Getenv(name)
	}
	// envs["FROM_USER"] = "997317653@qq.com"
	// envs["TO_USERS"] = "1781704348@qq.com"
	// envs["SECRET"] = "wyxyoqqzlhyzbbah"
	// envs["SMTP_SERVER"] = "smtp.qq.com"
	// envs["SMTP_PORT"] = "465"
	// envs["SUBJECT"] = "gomail"
	// envs["TEXT"] = "hello world"

	builder, err := NewBuilder(envs)
	if err != nil {
		fmt.Println("build faild: ", err)
		//os.Exit(1)
	}
	if err := builder.run(); err != nil {
		fmt.Println("build faild: ", err)
		//os.Exit(1)
	}
	fmt.Println("build success")
}

package main

import (
	"fmt"
	"os/exec"
	"strings"
	"io/ioutil"
)

// const baseSpace = "/root/src"

// Builder is
type Builder struct {
	Username    string
	Password         string
	Certificate        string
	Server          string
	Command string
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if envs["USERNAME"] == "" {
		return nil, fmt.Errorf("envionment variable USERNAME is required")
	}
	b.Username =  envs["USERNAME"]

	if envs["PASSWORD"] == "" {
		return nil, fmt.Errorf("envionment variable PASSWORD is required")
	}
	b.Password = envs["PASSWORD"]

	if envs["CERTIFICATE"] == "" {
		return nil, fmt.Errorf("envionment variable CERTIFICATE is required")
	}
	b.Certificate = envs["CERTIFICATE"]

	if envs["SERVER"] == "" {
		return nil, fmt.Errorf("envionment variable SERVER is required")
	}
	b.Server = envs["SERVER"]

	if envs["COMMAND"] == "" {
		return nil, fmt.Errorf("envionment variable COMMAND is required")
	}
	b.Command = envs["COMMAND"]

	return b, nil
}

func (b *Builder) run() error {
	//if err := os.Chdir(baseSpace); err != nil {
	//	return fmt.Errorf("Chdir to baseSpace(%s) failed:%v", baseSpace, err)
	//}

	if err := b.initConfig(); err != nil {
		return err
	}
	if err := b.runCommand(); err != nil {
		return err
	}

	return nil
}

func (b *Builder) initConfig() error {
	if err := ioutil.WriteFile("/root/cluster-ca.crt", []byte(b.Certificate), 0644); err != nil {
		fmt.Println("init config failed:", err)
		return err
	}

	commands := [][]string{
		// {"echo", os.Getenv( "CERTIFICATE"), ">", "/root/cluster-ca.crt"},
		// {"sh", "-c", "echo $CERTIFICATE > /root/cluster-ca.crt"},
		{"kubectl", "config", "set-credentials", "default-admin", fmt.Sprintf("--username=%s", b.Username), fmt.Sprintf("--password=%s", b.Password)},
		{"kubectl", "config", "set-cluster", "default-cluster", fmt.Sprintf("--server=%s", b.Server), "--certificate-authority=/root/cluster-ca.crt"},
		{"kubectl", "config", "set-context", "default-system", "--cluster=default-cluster", "--user=default-admin"},
		{"kubectl", "config", "use-context", "default-system"},
		//{"cat", "/root/cluster-ca.crt"},
		//{"cat", "/root/.kube/config"},
	}

	for _, command := range commands {
		if _, err := (CMD{Command: command}).Run(); err != nil {
			fmt.Println("init config failed:", err)
			return err
		}
	}

	return nil
}

func (b *Builder) runCommand() error {
	command := strings.Fields(b.Command)

	if command[0] != "kubectl" {
		command = append([]string{"kubectl"}, command...)
	}

	_, err := (CMD{Command: command}).Run()
	if err != nil {
		fmt.Println("run command failed:", err)
		return err
	}

	return nil
}


type CMD struct {
	Command []string // cmd with args
	WorkDir string
}

func (c CMD) Run() (string, error) {
	fmt.Println("Run CMD: ", strings.Join(c.Command, " "))

	cmd := exec.Command(c.Command[0], c.Command[1:]...)
	if c.WorkDir != "" {
		cmd.Dir = c.WorkDir
	}

	data, err := cmd.CombinedOutput()
	result := string(data)
	if len(result) > 0 {
		fmt.Println(result)
	}

	return result, err
}

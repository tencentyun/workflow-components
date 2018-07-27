package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

// Builder is
type Builder struct {
	Username    string
	Password    string
	Certificate string
	Server      string
	Cmd         string
	// Resource    string

	// output string
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if envs["USERNAME"] == "" {
		return nil, fmt.Errorf("envionment variable USERNAME is required")
	}
	b.Username = envs["USERNAME"]

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

	//if envs["RESOURCE"] == "" {
	//	return nil, fmt.Errorf("envionment variable RESOURCE is required")
	//}
	//b.Resource = envs["RESOURCE"]

	b.Cmd = envs["CMD"]
	return b, nil
}

func (b *Builder) run() error {
	if err := b.initConfig(); err != nil {
		return err
	}
	if err := b.execCmd(); err != nil {
		return err
	}
	// if err := b.RESOURCEapply(); err != nil {
	// 	return err
	// }

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

func (b *Builder) execCmd() error {
	//if b.Cmd == "" {
	//	b.output = b.Resource
	//}

	command := []string{"/bin/sh", "-c", b.Cmd}

	_, err := (CMD{Command: command}).Run()
	if err != nil {
		fmt.Println("exec CMD failed:", err)
		return err
	}

	// output = strings.TrimSpace(output)
	// if len(output) > 0 {
	// 	b.resourceYAML = output
	// } else {
	// 	return errors.New("output of CMD is empty")
	// }

	return nil
}

//func (b *Builder) apply() error {
//	if err := ioutil.WriteFile("/root/resource.yaml", []byte(b.output), 0644); err != nil {
//		fmt.Println("write resource yaml failed:", err)
//		return err
//	}
//
//	command := []string{"kubectl", "apply", "-f", "/root/resource.yaml"}
//
//	_, err := (CMD{Command: command}).Run()
//	if err != nil {
//		fmt.Println("run command failed:", err)
//		return err
//	}
//
//	return nil
//}

type CMD struct {
	Command []string // cmd with args
	WorkDir string
}

func (c CMD) Run() (string, error) {
	cmdStr := strings.Join(c.Command, " ")
	fmt.Printf("[%s] Run CMD: %s\n", time.Now().Format("2006-01-02 15:04:05"), cmdStr)

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

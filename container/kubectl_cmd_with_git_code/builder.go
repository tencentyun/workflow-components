package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
	"os"
	"path/filepath"
)

const baseSpace = "/root/src"

// Builder is
type Builder struct {
	GitCloneURL string
	GitRef      string

	Username    string
	Password    string
	Certificate string
	Server      string
	Cmd         string

	projectName string
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}


	if envs["GIT_CLONE_URL"] != "" {
		b.GitCloneURL = envs["GIT_CLONE_URL"]
		b.GitRef = envs["GIT_REF"]
	} else if envs["_WORKFLOW_GIT_CLONE_URL"] != "" {
		b.GitCloneURL = envs["_WORKFLOW_GIT_CLONE_URL"]
		b.GitRef = envs["_WORKFLOW_GIT_REF"]
	} else {
		return nil, fmt.Errorf("envionment variable GIT_CLONE_URL is required")
	}

	if b.GitRef == "" {
		b.GitRef = "master"
	}

	s := strings.TrimSuffix(strings.TrimSuffix(b.GitCloneURL, "/"), ".git")
	b.projectName = s[strings.LastIndex(s, "/")+1:]

	if b.GitRef = envs["GIT_REF"]; b.GitRef == "" {
		b.GitRef = "master"
	}

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

	b.Cmd = envs["CMD"]
	return b, nil
}

func (b *Builder) run() error {
	if err := b.gitPull(); err != nil {
		return err
	}

	if err := b.gitReset(); err != nil {
		return err
	}

	if err := b.initConfig(); err != nil {
		return err
	}
	if err := b.execCmd(); err != nil {
		return err
	}

	return nil
}

func (b *Builder) gitPull() error {
	var command = []string{"git", "clone", "--recurse-submodules", b.GitCloneURL, b.projectName}
	if _, err := (CMD{Command: command}).Run(); err != nil {
		fmt.Println("Clone project failed:", err)
		return err
	}
	fmt.Println("Clone project", b.GitCloneURL, "succeded.")
	return nil
}

func (b *Builder) gitReset() error {
	cwd, _ := os.Getwd()
	var command = []string{"git", "reset", "--hard", b.GitRef}
	if _, err := (CMD{command, filepath.Join(cwd, b.projectName)}).Run(); err != nil {
		fmt.Println("Switch to commit", b.GitRef, "failed:", err)
		return err
	}
	fmt.Println("Switch to", b.GitRef, "succeded.")
	return nil
}

func (b *Builder) initConfig() error {
	if err := ioutil.WriteFile("/root/cluster-ca.crt", []byte(b.Certificate), 0644); err != nil {
		fmt.Println("init config failed:", err)
		return err
	}

	commands := [][]string{
		{"kubectl", "config", "set-credentials", "default-admin", fmt.Sprintf("--username=%s", b.Username), fmt.Sprintf("--password=%s", b.Password)},
		{"kubectl", "config", "set-cluster", "default-cluster", fmt.Sprintf("--server=%s", b.Server), "--certificate-authority=/root/cluster-ca.crt"},
		{"kubectl", "config", "set-context", "default-system", "--cluster=default-cluster", "--user=default-admin"},
		{"kubectl", "config", "use-context", "default-system"},
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
	command := []string{"/bin/sh", "-c", b.Cmd}

	_, err := (CMD{Command: command}).Run()
	if err != nil {
		fmt.Println("exec CMD failed:", err)
		return err
	}

	return nil
}

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

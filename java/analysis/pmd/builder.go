package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Builder struct {
	GitCloneURL string
	GitRef      string
	projectName string
}

const baseSpace = "/root"

func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if envs["GIT_CLONE_URL"] != "" {
		b.GitCloneURL = envs["GIT_CLONE_URL"]
		if b.GitRef = envs["GIT_REF"]; b.GitRef == "" {
			b.GitRef = "master"
		}
	} else if envs["_WORKFLOW_GIT_CLONE_URL"] != "" {
		b.GitCloneURL = envs["_WORKFLOW_GIT_CLONE_URL"]
		b.GitRef = envs["_WORKFLOW_GIT_REF"]
	} else {
		return nil, fmt.Errorf("environment variable GIT_CLONE_URL is required")
	}

	s := strings.TrimSuffix(strings.TrimSuffix(b.GitCloneURL, "/"), ".git")
	b.projectName = s[strings.LastIndex(s, "/")+1:]

	return b, nil
}

func (b *Builder) run() error {
	if err := os.Chdir(baseSpace); err != nil {
		return fmt.Errorf("chdir to baseSpace(%s) failed:%v", baseSpace, err)
	}

	if err := b.gitPull(); err != nil {
		return err
	}

	if err := b.gitReset(); err != nil {
		return err
	}

	if err := b.build(); err != nil {
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
	fmt.Println("current: ", cwd)
	var command = []string{"git", "checkout", b.GitRef, "--"}
	if _, err := (CMD{command, filepath.Join(cwd, b.projectName)}).Run(); err != nil {
		fmt.Println("Switch to commit", b.GitRef, "failed:", err)
		return err
	}
	fmt.Println("Switch to", b.GitRef, "succeded.")

	return nil
}

func pathExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func (b *Builder) build() error {
	file := baseSpace + "/" + b.projectName + "/" + "build.gradle"
	if ok := pathExist(file); ok != true {
		return fmt.Errorf("file not exist")
	}

	var script = fmt.Sprintf("cat /root/pmd.conf >> %s", file)
	var command = []string{"sh", "-c", script}
	if _, err := (CMD{Command: command}).Run(); err != nil {
		fmt.Printf("Exec: build plugin.build failed: %v", err)
		return err
	}

	cwd, _ := os.Getwd()
	var command01 = []string{"gradle", "pmdMain"}
	(CMD{command01, filepath.Join(cwd, b.projectName)}).Run()

	var command02 = []string{"gradle", "pmdTest"}
	(CMD{command02, filepath.Join(cwd, b.projectName)}).Run()

	var mainFile = baseSpace + "/" + b.projectName + "/build/reports/pmd/main.xml"
	var testFile = baseSpace + "/" + b.projectName + "/build/reports/pmd/test.xml"
	var mainCommand = []string{"cat", mainFile}
	var testCommand = []string{"cat", testFile}
	(CMD{Command: mainCommand}).Run()
	(CMD{Command: testCommand}).Run()

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

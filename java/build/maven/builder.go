package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const baseSpace = "/root/src"

type Builder struct {
	// 用户提供参数, 通过环境变量传入
	GitCloneURL string
	GitRef      string
	Goals       string
	PomPath     string

	projectName string
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if envs["GIT_CLONE_URL"] == "" {
		return nil, fmt.Errorf("envionment variables GIT_CLONE_URL is required")
	}

	b.GitCloneURL = envs["GIT_CLONE_URL"]
	s := strings.TrimSuffix(strings.TrimSuffix(b.GitCloneURL, "/"), ".git")
	b.projectName = s[strings.LastIndex(s, "/")+1:]

	if b.GitRef = envs["GIT_REF"]; b.GitRef == "" {
		b.GitRef = "master"
	}

	b.Goals = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(envs["GOALS"]), "mvn "))

	if b.PomPath = envs["POM_PATH"]; b.PomPath == "" {
		b.PomPath = "./pom.xml"
	}

	return b, nil
}

func (b *Builder) run() error {
	if err := os.Chdir(baseSpace); err != nil {
		return fmt.Errorf("Chdir to baseSpace(%s) failed:%v", baseSpace, err)
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

	if err := b.handleArtifact(); err != nil {
		return err
	}
	// err = b.doPush(b.Image)
	// if err != nil {
	// 	return
	// }
	return nil
}

func (b *Builder) build() error {
	var command = strings.Fields(b.Goals)

	if len(command) == 0 {
		command = append(command, "mvn", "package")
	}

	if command[0] != "mvn" {
		command = append([]string{"mvn"}, command...)
	}

	command = append(command, "-f", b.PomPath)

	cwd, _ := os.Getwd()
	if _, err := (CMD{command, filepath.Join(cwd, b.projectName)}).Run(); err != nil {
		fmt.Println("Run mvn goals failed:", err)
		return err
	}
	fmt.Println("Run mvn goals succeded.")
	return nil
}

func (b *Builder) handleArtifact() error {
	command := []string{"find", "./target", "-name", "*.jar", "-o", "-name", "*.war"}

	cwd, _ := os.Getwd()

	output, err := (CMD{command, filepath.Join(cwd, b.projectName)}).Run()
	if err != nil {
		fmt.Println("Run find artifact failed:", err)
		return err
	}

	output = strings.TrimSpace(output)
	if len(output) > 0 {
		artifact := strings.Join(strings.Split(output, "\n"), ";")
		fmt.Printf("[JOB_OUT] ARTIFACT = %s\n", artifact)
	} else {
		return errors.New("no artifact")
	}

	fmt.Println("Run handle artifact succeded.")
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

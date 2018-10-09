package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	baseSpace  = "/root/src"
	cacheSpace = "/workflow-cache"
)

type Builder struct {
	// 用户提供参数, 通过环境变量传入
	GitCloneURL     string
	GitRef          string
	AnalysisOptions string
	AnalysisTarget  string

	projectName string

	WorkflowCache bool
	workDir       string
	gitDir        string
}

// NewBuilder is
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

	b.AnalysisOptions = envs["ANALYSIS_OPTIONS"]
	if b.AnalysisTarget = envs["ANALYSIS_TARGET"]; b.AnalysisTarget == "" {
		b.AnalysisTarget = ""
	}

	s := strings.TrimSuffix(strings.TrimSuffix(b.GitCloneURL, "/"), ".git")
	b.projectName = s[strings.LastIndex(s, "/")+1:]

	b.WorkflowCache = strings.ToLower(envs["_WORKFLOW_FLAG_CACHE"]) == "true"

	if b.WorkflowCache {
		b.workDir = cacheSpace
	} else {
		b.workDir = baseSpace
	}
	b.gitDir = filepath.Join(b.workDir, b.projectName)

	return b, nil
}

func (b *Builder) run() error {
	if err := os.Chdir(b.workDir); err != nil {
		return fmt.Errorf("chdir to workdir (%s) failed:%v", b.workDir, err)
	}

	if _, err := os.Stat(b.gitDir); os.IsNotExist(err) {
		if err := b.gitPull(); err != nil {
			return err
		}

		if err := b.gitReset(); err != nil {
			return err
		}
	}

	if err := b.build(); err != nil {
		return err
	}

	return nil
}

func (b *Builder) build() error {
	var command = []string{"java", "-jar", "/root/src/checkstyle-8.10-all.jar"}

	if b.AnalysisOptions != "" {
		command = append(command, b.AnalysisOptions)
	}
	command = append(command, "-c")
	command = append(command, "/root/src/checkstyle.xml")
	command = append(command, filepath.Join(b.gitDir, b.AnalysisTarget))

	if _, err := (CMD{Command: command}).Run(); err != nil {
		fmt.Println("Run checkstyle failed:", err)
		return err
	}
	fmt.Println("Run checkstyle succeded.")
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
	var command = []string{"git", "checkout", b.GitRef, "--"}
	if _, err := (CMD{command, b.gitDir}).Run(); err != nil {
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

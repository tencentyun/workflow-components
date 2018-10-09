package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Builder struct {
	GitCloneURL string
	GitRef      string
	projectName string
	OutputXml   bool

	WorkflowCache bool
	workDir       string
	gitDir        string
}

const (
	baseSpace  = "/root/src"
	cacheSpace = "/workflow-cache"
)

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
	b.WorkflowCache = strings.ToLower(envs["_WORKFLOW_FLAG_CACHE"]) == "true"

	if b.WorkflowCache {
		b.workDir = cacheSpace
	} else {
		b.workDir = baseSpace
	}
	b.gitDir = filepath.Join(b.workDir, b.projectName)

	b.OutputXml = strings.ToLower(envs["OUT_PUT_XML"]) == "true"

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

	if err := b.preBuild(); err != nil {
		return err
	}

	if err := b.build(); err != nil {
		return err
	}

	if b.OutputXml == true {
		if err := b.afterBuild(); err != nil {
			return err
		}
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
	if _, err := (CMD{command, b.gitDir}).Run(); err != nil {
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

func (b *Builder) preBuild() error {
	file := b.gitDir + "/" + "build.gradle"
	if ok := pathExist(file); ok != true {
		return fmt.Errorf("build.gradle file not exist")
	}

	var pmd = "echo gradle -q tasks --all | grep pmd"
	var pmdIsExistCommand = []string{"sh", "-c", pmd}
	if isExist, _ := (CMD{Command: pmdIsExistCommand}).Run(); isExist == "" {
		return fmt.Errorf("gradle pmd not exist")
	}
	return nil
}

func (b *Builder) build() error {
	var command01 = []string{"gradle", "pmdMain"}
	(CMD{command01, b.gitDir}).Run()

	var command02 = []string{"gradle", "pmdTest"}
	(CMD{command02, b.gitDir}).Run()

	return nil
}

func showXmlReport(file string) error {
	fmt.Printf("SHOW REPORT: %s", file)
	inputFile, inputError := os.Open(file)
	if inputError != nil {
		return fmt.Errorf("the file: %s not exist\n", file)
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("%s", inputString)
		if readerError == io.EOF {
			return nil
		}
	}
}

func (b *Builder) afterBuild() error {
	var mainFile = b.gitDir + "/build/reports/pmd/main.xml"
	var testFile = b.gitDir + "/build/reports/pmd/test.xml"
	if err := showXmlReport(mainFile); err != nil {
		return err
	}
	if err := showXmlReport(testFile); err != nil {
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

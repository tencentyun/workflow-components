package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	baseSpace  = "/go/src"
	cacheSpace = "/workflow-cache"
)

type Builder struct {
	GitCloneURL        string
	GitRef             string
	GtestPackageOrFile string
	Output             string
	BuildPackageName   string
	PackageTarget      string
	Parent             string
	ProjectName        string
	BuildVendor        string

	WorkflowCache bool

	workDir string
	gitDir  string
}

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

	b.Output = envs["OUTPUT"]
	b.PackageTarget = envs["PACKAGE_TARGET"]
	b.BuildVendor = envs["BUILD_VENDOR_CMD"]

	b.WorkflowCache = strings.ToLower(envs["_WORKFLOW_FLAG_CACHE"]) == "true"

	// 设置工作目录主要用于设置gopath, gobin
	// 设置正确的工作目录，主要有三个，workDir, parent, gitDit
	if b.WorkflowCache {
		b.workDir = fmt.Sprintf("%s/src", cacheSpace)
	} else {
		b.workDir = baseSpace
	}
	// 设置git path
	b.BuildPackageName = envs["BUILD_PACKAGE_NAME"]
	// TODO 对go pack进行处理
	b.gitDir = fmt.Sprintf("%s/%s", b.workDir, b.BuildPackageName)
	// fmt.Println(b.gitDir)

	parent, err := CreateGoPackageParentPath(b.gitDir)
	if err != nil {
		return nil, err
	}
	b.Parent = parent
	b.ProjectName = b.BuildPackageName[strings.LastIndex(b.BuildPackageName, "/")+1:]

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

	if b.BuildVendor != "" {
		if err := b.preBuild(); err != nil {
			return err
		}
	}

	if err := b.build(); err != nil {
		return err
	}
	return nil
}

func (b *Builder) gitPull() error {
	var command = []string{"git", "clone", "--recurse-submodules", b.GitCloneURL, b.ProjectName}
	if _, err := (CMD{Command: command, WorkDir: b.Parent}).Run(); err != nil {
		fmt.Println("Clone project failed:", err)
		return err
	}
	fmt.Println("Clone project", b.GitCloneURL, "succeed.")
	return nil
}

func (b *Builder) gitReset() error {
	cwd, _ := os.Getwd()
	fmt.Println("current: ", cwd)
	var command = []string{"git", "checkout", b.GitRef, "--"}
	if _, err := (CMD{command, b.gitDir, nil}).Run(); err != nil {
		fmt.Println("Switch to commit", b.GitRef, "failed:", err)
		return err
	}
	fmt.Println("Switch to", b.GitRef, "succeed.")

	return nil
}

func (b *Builder) preBuild() error {
	var env []string
	if b.WorkflowCache {
		env = []string{"GOPATH=/go:/workflow-cache", "PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"}
	}
	var command = []string{"/bin/sh", "-c", b.BuildVendor}
	// var command = []string{"/bin/sh", "-c", fmt.Sprintf("%q", b.BuildVendor)}
	if _, err := (CMD{command, b.gitDir, env}).Run(); err != nil {
		fmt.Println("Run golang build failed:", err)
		return err
	}
	fmt.Println("Run golang build succeed.")
	return nil
}

func (b *Builder) build() error {
	var env []string
	if b.WorkflowCache {
		env = []string{"GOPATH=/go:/workflow-cache", "PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"}
	}
	var command = []string{"go", "build", "-v"}
	if b.Output != "" {
		command = append(command, "-o", b.Output)
	}
	// if b.PackageTarget != "." {
	command = append(command, b.PackageTarget)
	// }

	if _, err := (CMD{command, b.gitDir, env}).Run(); err != nil {
		fmt.Println("Run golang build failed:", err)
		return err
	}
	fmt.Println("Run golang build succeed.")
	return nil
}

type CMD struct {
	Command []string // cmd with args
	WorkDir string
	Env     []string
}

func (c CMD) Run() (string, error) {
	fmt.Println("Run CMD: ", strings.Join(c.Command, " "))

	cmd := exec.Command(c.Command[0], c.Command[1:]...)

	if c.WorkDir != "" {
		cmd.Dir = c.WorkDir
	}
	if c.Env != nil {
		cmd.Env = c.Env
	}

	data, err := cmd.CombinedOutput()
	result := string(data)
	if len(result) > 0 {
		fmt.Println(result)
	}

	return result, err
}

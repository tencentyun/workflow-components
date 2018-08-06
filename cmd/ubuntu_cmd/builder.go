package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const baseSpace = "/root/src"

// Builder is
type Builder struct {
	Cmd string

	projectName string
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	b.Cmd = envs["CMD"]
	return b, nil
}

func (b *Builder) run() error {
	if err := b.execCmd(); err != nil {
		return err
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

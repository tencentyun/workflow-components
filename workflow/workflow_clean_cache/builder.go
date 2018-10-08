package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const baseSpace = "/root/src"

// Builder is
type Builder struct {
}

// NewBuilder is
func NewBuilder() (*Builder, error) {
	b := &Builder{}

	return b, nil
}

func (b *Builder) run() error {
	if err := b.cleanCache(); err != nil {
		return err
	}

	return nil
}

func (b *Builder) cleanCache() error {
	// command := []string{"rm", "-rf", "/workflow-cache/*"} // why this not working?
	command := []string{"/bin/sh", "-c", "rm -rf /workflow-cache/*"}
	if _, err := (CMD{Command: command}).Run(); err != nil {
		fmt.Println("clean cache failed:", err)
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

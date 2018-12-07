package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const baseSpace = "/root/src"

// Builder is
type Builder struct {
	CacheID string
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{
		CacheID: envs["CACHE_ID"],
	}

	if b.CacheID == "" {
		return nil, fmt.Errorf("envionment variable CACHE_ID is required")
	}

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
	(CMD{Command: []string{"/bin/sh", "-c", "ls /workflow-cache/"}}).Run()
	command := []string{"/bin/sh", "-c", fmt.Sprintf("rm -rf /workflow-cache/%s", b.CacheID)}
	if _, err := (CMD{Command: command}).Run(); err != nil {
		fmt.Println("clean cache failed:", err)
		return err
	}
	(CMD{Command: []string{"/bin/sh", "-c", "ls /workflow-cache/"}}).Run()

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

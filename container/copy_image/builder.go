package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const baseSpace = "/root/src"

// Builder is
type Builder struct {
	// 用户提供参数, 通过环境变量传入
	Image          string
	HubUser  string
	HubToken string

	ToImage          string
	ToHubUser  string
	ToHubToken string

	imageTag string
	hub           string
	toHub           string
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	if envs["IMAGE_TAG"] == "" {
		return nil, fmt.Errorf("envionment variable IMAGE_TAG is required")
	}
	b.Image = envs["IMAGE_TAG"]

	if imageAndTag := strings.Split(b.Image, ":"); len(imageAndTag) > 1 {
		b.imageTag = imageAndTag[len(imageAndTag)-1]
	} else {
		b.imageTag = "latest"
	}

	b.HubUser = envs["HUB_USER"]
	b.HubToken = envs["HUB_TOKEN"]
	if b.HubUser == "" && b.HubToken == "" {
		b.HubUser = envs["_WORKFLOW_HUB_USER"]
		b.HubToken = envs["_WORKFLOW_HUB_TOKEN"]
	}
	if b.HubUser == "" || b.HubToken == "" {
		return nil, fmt.Errorf("envionment variable HUB_USER, HUB_TOKEN are required")
	}

	if envs["TO_IMAGE"] == "" {
		return nil, fmt.Errorf("envionment variable TO_IMAGE is required")
	}
	b.ToImage = envs["TO_IMAGE"]

	if imageAndTag := strings.Split(b.ToImage, ":"); len(imageAndTag) <= 1 {
		b.ToImage = fmt.Sprintf("%s:%s", b.ToImage, b.imageTag)
	}


	b.ToHubUser = envs["TO_HUB_USER"]
	b.ToHubToken = envs["TO_HUB_TOKEN"]
	if b.ToHubUser == "" || b.ToHubToken == "" {
		return nil, fmt.Errorf("envionment variable TO_HUB_USER, TO_HUB_TOKEN are required")
	}

	if strings.Index(b.Image, ".") > -1 {
		b.hub = b.Image
	} else {
		b.hub = "index.docker.io" // default server
	}

	if strings.Index(b.ToImage, ".") > -1 {
		b.toHub = b.ToImage
	} else {
		b.toHub = "index.docker.io" // default server
	}

	return b, nil
}

func (b *Builder) run() error {
	if err := os.Chdir(baseSpace); err != nil {
		return fmt.Errorf("Chdir to baseSpace(%s) failed:%v", baseSpace, err)
	}

	if err := b.loginRegistry(b.hub, b.HubUser, b.HubToken); err != nil {
		return err
	}

	if err := b.pull(b.Image); err != nil {
		return err
	}

	if err := b.tag(b.Image, b.ToImage); err != nil {
		return err
	}

	if err := b.loginRegistry(b.toHub, b.ToHubUser, b.ToHubToken); err != nil {
		return err
	}

	if err := b.push(b.ToImage); err != nil {
		return err
	}

	if err := b.pluckImageID(b.ToImage); err != nil {
		return err
	}

	if err := b.pluckImageDigest(b.ToImage); err != nil {
		return err
	}

	if err := b.cleanImage(b.Image); err != nil {
		return err
	}
	if err := b.cleanImage(b.ToImage); err != nil {
		return err
	}
	return nil
}


func (b *Builder) pull(imageURL string) error {
	var command = []string{"docker", "pull", imageURL}
	if _, err := (CMD{Command: command}).Run(); err != nil {
		fmt.Println("Run docker pull failed:", err)
		return err
	}
	fmt.Println("Run docker pull succeed.")
	return nil
}

func (b *Builder) push(imageURL string) error {
	var command = []string{"docker", "push", imageURL}
	if _, err := (CMD{Command: command}).Run(); err != nil {
		fmt.Println("Run docker push failed:", err)
		return err
	}
	fmt.Println("Run docker push succeed.")
	return nil
}

func (b *Builder) tag(oldImage, newImage string) error {
	var command = []string{"docker", "tag", oldImage, newImage}
	if _, err := (CMD{Command: command}).Run(); err != nil {
		fmt.Println("Run docker tag failed:", err)
		return err
	}
	fmt.Println("Run docker tag succeed.")
	return nil
}



func (b *Builder) loginRegistry(hub, user, token string) error {
	var command = []string{"docker", "login", hub, "-u", user, "-p", token}
	if _, err := (CMD{Command: command}).Run(); err != nil {
		fmt.Println("docker login failed:", err)
		return err
	}
	fmt.Println("docker login succ.")
	return nil
}


func (b *Builder) pluckImageID(imageURL string) error {
	// docker inspect hub.cloud.tencent.com/tencenthub/docker_builder:latest --format '{{.Id}}'
	var command = []string{"docker", "inspect", imageURL, "--format", "{{.Id}}"}
	// docker images ccr.ccs.tencentyun.com/tencenthub/workflow:latest --format "{{.ID}}"
	// var command = []string{"docker", "images", b.Image, "--format", "{{.ID}}"}
	output, err := (CMD{Command: command}).Run()

	if err != nil {
		fmt.Println("pluck image id failed:", err)
		return err
	}
	if len(output) > 0 {
		fmt.Println("pluck image id succeed.")
		fmt.Printf("[JOB_OUT] IMAGE_ID = %s", output)
	} else {
		return errors.New("Can not get image id")
	}

	return nil
}

func (b *Builder) pluckImageDigest(imageURL string) error {
	// docker inspect hub.cloud.tencent.com/tencenthub/docker_builder:latest --format '{{index .RepoDigests 0}}'
	var command = []string{"docker", "inspect", imageURL, "--format", "{{index .RepoDigests 0}}"}
	output, err := (CMD{Command: command}).Run()

	if err != nil {
		fmt.Println("pluck image digest failed:", err)
		return err
	}
	cut := b.Image + "@"
	output = strings.TrimPrefix(output, cut)
	if len(output) > 0 {
		fmt.Println("pluck image digest succeed.")
		fmt.Printf("[JOB_OUT] IMAGE_DIGEST = %s\n", output)
	} else {
		return errors.New("Can not get image digest")
	}

	return nil
}

func (b *Builder) cleanImage(imageURL string) error {
	var command = []string{"docker", "rmi", imageURL}
	if _, err := (CMD{Command: command}).Run(); err != nil {
		fmt.Println("Run docker rmi", imageURL, "failed:", err)
		return err
	}
	fmt.Println("clean local image completely.")
	return nil
}

func ensureDirExists(dir string) (err error) {
	f, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(dir, os.FileMode(0755))
		}
		return err
	}

	if !f.IsDir() {
		return fmt.Errorf("%s is not dir", dir)
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

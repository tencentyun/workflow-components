package builder

import (
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	"fmt"
)

type DeployInfo struct {
	Strategy    string
	Template    *extv1beta1.Deployment
	DeployGroup string
	Services    []string
	Image       string
	Replicas    int32
	// ServicesGate bool

	Version *int64
	Name    string
}

func (info *DeployInfo) String() string {
	return fmt.Sprintf("Template: %s,Strategy: %s, DeployGroup: %s, Services: %s, Image: %s, Replicas: %d, Version: %d, Name: %s",
		info.Template.Name, info.Strategy, info.DeployGroup, info.Services, info.Image, info.Replicas, info.Version, info.Name)
}

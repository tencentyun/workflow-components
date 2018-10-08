package main

import (
	"kubecd/k8s"
	"errors"
	"fmt"
	"io/ioutil"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	"path/filepath"
	"strconv"
	"strings"
)

//const baseSpace = "/Users/testuser"
const baseSpace = "/root"

// Builder is
type Builder struct {
	Action string

	// 用户提供参数, 通过环境变量传入
	Username    string
	Password    string
	Certificate string
	Server      string
	Namespace   string

	Services []string
	Replicas int32
	Strategy string
	Image    string

	DeployTarget   string
	DeploymentName string

	// rollback
	FromDeployTarget   string
	FromDeploymentName string
	ToDeployTarget     string
	ToDeploymentName   string

	DeployGroup string

	// ScaleStep int64
	ScaleTo   *int32
	ScaleUp   *int32
	ScaleDown *int32

	ShrinkTo     int64
	AutoDeletion bool

	cluster *k8s.Cluster
}

// NewBuilder is
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}
	if envs["USERNAME"] == "" {
		return nil, fmt.Errorf("environment variable USERNAME is required")
	}
	b.Username = envs["USERNAME"]

	if envs["PASSWORD"] == "" {
		return nil, fmt.Errorf("environment variable PASSWORD is required")
	}
	b.Password = envs["PASSWORD"]

	if envs["CERTIFICATE"] == "" {
		return nil, fmt.Errorf("environment variable CERTIFICATE is required")
	}
	b.Certificate = envs["CERTIFICATE"]

	if envs["SERVER"] == "" {
		return nil, fmt.Errorf("environment variable SERVER is required")
	}
	b.Server = envs["SERVER"]

	b.Namespace = envs["NAMESPACE"]
	if b.Namespace == "" {
		b.Namespace = "default"
	}

	if envs["ACTION"] == "" {
		return nil, fmt.Errorf("environment variable ACTION is required")
	}

	b.Action = envs["ACTION"]
	switch b.Action {
	case "deploy":
		if envs["IMAGE"] == "" {
			return nil, fmt.Errorf("environment variable IMAGE is required")
		}
		b.Image = envs["IMAGE"]

		if envs["DEPLOY_GROUP"] == "" {
			return nil, fmt.Errorf("environment variable DEPLOY_GROUP is required")
		}
		b.DeployGroup = envs["DEPLOY_GROUP"]

		if envs["REPLICAS"] == "" {
			envs["REPLICAS"] = "0"
		}
		replicas, err := strconv.ParseInt(envs["REPLICAS"], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid environment variable REPLICAS: %s", envs["REPLICAS"])
		}
		b.Replicas = int32(replicas)

		if envs["STRATEGY"] != k8s.StrategyRecreate && envs["STRATEGY"] != k8s.StrategyBlueGreen &&
			envs["STRATEGY"] != k8s.StrategyCanary && envs["STRATEGY"] != k8s.StrategyOffline {
			return nil, fmt.Errorf("invalid environment variable STRATEGY: %s", envs["STRATEGY"])
		}

		b.Strategy = envs["STRATEGY"]

		if envs["SERVICES"] != "" {
			b.Services = strings.Split(envs["SERVICES"], ",")
		}

		b.DeploymentName = envs["DEPLOYMENT_NAME"]
		b.DeployTarget = envs["DEPLOY_TARGET"]

	case "scale":
		if envs["DEPLOYMENT_NAME"] != "" {
			b.DeploymentName = envs["DEPLOYMENT_NAME"]
		} else if envs["DEPLOY_TARGET"] != "" && envs["DEPLOY_GROUP"] != "" {
			b.DeployTarget = envs["DEPLOY_TARGET"]
			b.DeployGroup = envs["DEPLOY_GROUP"]
		} else {
			return nil, errors.New("environment variable DEPLOYMENT_NAME or (DEPLOY_TARGET, DEPLOY_GROUP) is required")
		}

		var scale int64
		var err error
		if envs["SCALE_TO"] != "" {
			scale, err = strconv.ParseInt(envs["SCALE_TO"], 10, 64)
			s32 := int32(scale)
			b.ScaleTo = &s32
		} else if envs["SCALE_UP"] != "" {
			scale, err = strconv.ParseInt(envs["SCALE_UP"], 10, 64)
			s32 := int32(scale)
			b.ScaleUp = &s32
		} else if envs["SCALE_DOWN"] != "" {
			scale, err = strconv.ParseInt(envs["SCALE_DOWN"], 10, 64)
			s32 := int32(scale)
			b.ScaleDown = &s32
		} else {
			return nil, errors.New("environment variable SCALE_TO or SCALE_UP or SCALE_DOWN is required")
		}
		if err != nil {
			return nil, err
		}

		b.AutoDeletion = strings.ToLower(envs["AUTO_DELETION"]) == "true"

	case "disable", "enable":
		if envs["DEPLOYMENT_NAME"] != "" {
			b.DeploymentName = envs["DEPLOYMENT_NAME"]
		} else if envs["DEPLOY_TARGET"] != "" && envs["DEPLOY_GROUP"] != "" {
			b.DeployTarget = envs["DEPLOY_TARGET"]
			b.DeployGroup = envs["DEPLOY_GROUP"]
		} else {
			return nil, errors.New("environment variable DEPLOYMENT_NAME or (DEPLOY_TARGET, DEPLOY_GROUP) is required")
		}

	case "delete":
		if envs["DEPLOYMENT_NAME"] != "" {
			b.DeploymentName = envs["DEPLOYMENT_NAME"]
		} else if envs["DEPLOY_TARGET"] != "" && envs["DEPLOY_GROUP"] != "" {
			b.DeployTarget = envs["DEPLOY_TARGET"]
			b.DeployGroup = envs["DEPLOY_GROUP"]
		} else {
			return nil, errors.New("environment variable DEPLOYMENT_NAME or (DEPLOY_TARGET, DEPLOY_GROUP) is required")
		}

	case "shrink":
		if envs["DEPLOY_GROUP"] == "" {
			return nil, fmt.Errorf("environment variable DEPLOY_GROUP is required")
		}
		b.DeployGroup = envs["DEPLOY_GROUP"]

		if envs["SHRINK_TO"] == "" {
			return nil, fmt.Errorf("environment variable SHRINK_TO is required")
		}
		shrinkTo, err := strconv.ParseInt(envs["SHRINK_TO"], 10, 64)
		if err != nil {
			return nil, err
		}
		b.ShrinkTo = shrinkTo

	case "rollback":
		if envs["FROM_DEPLOYMENT_NAME"] != "" {
			b.FromDeploymentName = envs["FROM_DEPLOYMENT_NAME"]
		} else if envs["FROM_DEPLOY_TARGET"] != "" && envs["DEPLOY_GROUP"] != "" {
			b.FromDeployTarget = envs["FROM_DEPLOY_TARGET"]
			b.DeployGroup = envs["DEPLOY_GROUP"]
		} else {
			return nil, fmt.Errorf("environment variable FROM_DEPLOYMENT_NAME or (FROM_DEPLOY_TARGET, DEPLOY_GROUP) is required")
		}

		if envs["TO_DEPLOYMENT_NAME"] != "" {
			b.ToDeploymentName = envs["TO_DEPLOYMENT_NAME"]
		} else if envs["TO_DEPLOY_TARGET"] != "" && envs["DEPLOY_GROUP"] != "" {
			b.ToDeployTarget = envs["TO_DEPLOY_TARGET"]
			b.DeployGroup = envs["DEPLOY_GROUP"]
		} else {
			return nil, fmt.Errorf("environment variable TO_DEPLOYMENT_NAME or (TO_DEPLOY_TARGET, DEPLOY_GROUP) is required")
		}

	default:
		return nil, fmt.Errorf("invalid action %s", b.Action)
	}
	return b, nil
}

func (b *Builder) run() error {
	err := b.initConfig()
	if err != nil {
		return err
	}

	switch b.Action {
	case "deploy":
		return b.deploy()
	case "scale":
		return b.scale()
	case "disable":
		return b.disable()
	case "enable":
		return b.enable()
	case "delete":
		return b.delete()
	case "shrink":
		return b.shrink()
	case "rollback":
		return b.rollback()
	default:
		return fmt.Errorf("invalid action %s", b.Action)
	}
	// pp.Print(newDm)

	return nil
}

func (b *Builder) deploy() (err error) {
	c := b.cluster

	var dm *extv1beta1.Deployment // 模板
	if b.DeploymentName != "" {
		dm, err = c.GetDeployment(b.DeploymentName)
	} else {
		target := b.DeployTarget
		if target == "" {
			target = "newest"
		}
		dm, err = c.GetDeploymentByDeployGroup(b.DeployGroup, target)
	}
	if err != nil {
		return
	}

	info := &k8s.DeployInfo{
		Strategy:    b.Strategy,
		Template:    dm,
		DeployGroup: b.DeployGroup,
		Services:    b.Services,
		Image:       b.Image,
		Replicas:    b.Replicas,
	}

	newDm, err := c.Deploy(info)
	if err != nil {
		return err
	}

	version, err := c.GetDeploymentVersion(newDm)
	if err != nil {
		return err
	}
	fmt.Printf("[JOB_OUT] NEW_DEPLOY_NAME=%s\n", newDm.Name)
	fmt.Printf("[JOB_OUT] NEW_DEPLOY_VERSION=%d\n", version)
	return nil
}

//func (b *Builder) deployFromService() error {
//	if b.Services == "" {
//		return errors.New("environment variable SERVICES is required")
//	}
//	if b.Image == "" {
//		return fmt.Errorf("environment variable IMAGE is required")
//	}
//	c := b.cluster
//	svc, err := c.GetService(b.Services)
//	if err != nil {
//		return err
//	}
//
//	dmList, err := c.ListEnabledDeploymentsByService(svc)
//
//	if err != nil {
//		return err
//	}
//
//	dm, err :=  c.Deploy(dmList, b.Strategy, b.Services, b.Image, b.Replicas)
//	if err != nil {
//		return err
//	}
//
//	version, err := c.GetDeploymentVersion(dm)
//	if err != nil {
//		return err
//	}
//	fmt.Printf("[JOB_OUT] NEW_DEPLOY_NAME=%s\n", dm.Name)
//	fmt.Printf("[JOB_OUT] NEW_DEPLOY_VERSION=%d\n", version)
//	return nil
//}

func (b *Builder) scale() (err error) {
	c := b.cluster

	var dm *extv1beta1.Deployment
	if b.DeploymentName != "" {
		dm, err = c.GetDeployment(b.DeploymentName)
	} else {
		dm, err = c.GetDeploymentByDeployGroup(b.DeployGroup, b.DeployTarget)
	}
	if err != nil {
		return
	}

	//svc, err := c.GetService(b.Services)
	//if err != nil {
	//	return err
	//}

	var scaleTo int32
	if b.ScaleTo != nil {
		scaleTo = *b.ScaleTo
	} else if b.ScaleUp != nil {
		scaleTo = *dm.Spec.Replicas + *b.ScaleUp
	} else if b.ScaleDown != nil {
		scaleTo = *dm.Spec.Replicas - *b.ScaleDown
	}

	if err := c.ScaleDeployment(dm, scaleTo, b.AutoDeletion); err != nil {
		return err
	}

	fmt.Printf("[JOB_OUT] SCALED_DEPLOY_NAME=%s\n", dm.Name)
	fmt.Printf("[JOB_OUT] SCALED_DEPLOY_REPLICAS=%d\n", scaleTo)
	return nil
}

func (b *Builder) disable() (err error) {
	c := b.cluster
	var dm *extv1beta1.Deployment
	if b.DeploymentName != "" {
		dm, err = c.GetDeployment(b.DeploymentName)
	} else {
		dm, err = c.GetDeploymentByDeployGroup(b.DeployGroup, b.DeployTarget)
	}
	if err != nil {
		return
	}

	if err := c.DisableDeploymentAllTraffic(dm); err != nil {
		return err
	}

	return nil
}

func (b *Builder) enable() (err error) {
	c := b.cluster
	var dm *extv1beta1.Deployment
	if b.DeploymentName != "" {
		dm, err = c.GetDeployment(b.DeploymentName)
	} else {
		dm, err = c.GetDeploymentByDeployGroup(b.DeployGroup, b.DeployTarget)
	}
	if err != nil {
		return
	}

	if err := c.EnableDeploymentAllTraffic(dm); err != nil {
		return err
	}

	return nil
}

func (b *Builder) delete() (err error) {
	c := b.cluster
	var dm *extv1beta1.Deployment
	if b.DeploymentName != "" {
		dm, err = c.GetDeployment(b.DeploymentName)
	} else {
		dm, err = c.GetDeploymentByDeployGroup(b.DeployGroup, b.DeployTarget)
	}
	if err != nil {
		return
	}
	// envs := b.envs
	// svc, err := c.GetService(b.Services)
	// if err != nil {
	// 	return err
	// }
	//
	// dm, err := c.GetDeploymentByTarget(svc, b.DeployTarget)
	// if err != nil {
	// 	return err
	// }

	if err := c.DeleteDeployment(dm); err != nil {
		return err
	}

	fmt.Printf("[JOB_OUT] DELETED_DEPLOY_NAME=%s\n", dm.Name)
	version, _ := c.GetDeploymentVersion(dm)
	fmt.Printf("[JOB_OUT] DELETED_DEPLOY_VERSION=%d\n", version)
	return nil
}

// TODO 流量版本删完的提醒
// TODO 提供一个参数: 只删除无流量的, 或者优先删除无流量的
func (b *Builder) shrink() (err error) {
	c := b.cluster

	dmList, err := c.ListAllDeploymentsByDeployGroup(b.DeployGroup)
	if err != nil {
		return err
	}

	deletionLength := len(dmList) - int(b.ShrinkTo)

	if deletionLength <= 0 {
		fmt.Printf("[JOB_OUT] DELETION_LENGTH=%d\n", 0)
		return nil
	}

	for i := 0; i < deletionLength; i++ {
		if err := c.ScaleDeployment(dmList[i], 0, true); err != nil {
			return err
		}
	}

	fmt.Printf("[JOB_OUT] DELETION_LENGTH=%d\n", deletionLength)
	return nil
}

//
func (b *Builder) rollback() (err error) {
	c := b.cluster

	var fromDm, toDm *extv1beta1.Deployment

	if b.FromDeploymentName != "" {
		fromDm, err = c.GetDeployment(b.FromDeploymentName)
	} else {
		fromDm, err = c.GetDeploymentByDeployGroup(b.DeployGroup, b.FromDeployTarget)
	}
	if err != nil {
		return
	}
	if fromDm == nil {
		return errors.New("cat not find from deployment")
	}

	if b.ToDeploymentName != "" {
		toDm, err = c.GetDeployment(b.ToDeploymentName)
	} else {
		toDm, err = c.GetDeploymentByDeployGroup(b.DeployGroup, b.ToDeployTarget)
	}
	if err != nil {
		return
	}
	if toDm == nil {
		return errors.New("cat not find to deployment")
	}

	if fromDm.Name == toDm.Name {
		return errors.New("from deployment is the same as to deployment")
	}

	if err = c.RollbackDeployment(fromDm, toDm); err != nil {
		return err
	}

	return nil
}

func (b *Builder) initConfig() error {
	ca := filepath.Join(baseSpace, "cluster-ca.crt")
	if err := ioutil.WriteFile(ca, []byte(b.Certificate), 0644); err != nil {
		fmt.Println("init config failed:", err)
		return err
	}

	commands := [][]string{
		{"kubectl", "config", "set-credentials", "default-admin", fmt.Sprintf("--username=%s", b.Username), fmt.Sprintf("--password=%s", b.Password)},
		{"kubectl", "config", "set-cluster", "default-cluster", fmt.Sprintf("--server=%s", b.Server), fmt.Sprintf("--certificate-authority=%s", ca)},
		{"kubectl", "config", "set-context", "default-system", "--cluster=default-cluster", "--user=default-admin"},
		{"kubectl", "config", "use-context", "default-system"},
	}

	for _, command := range commands {
		if _, err := (CMD{Command: command}).Run(); err != nil {
			fmt.Println("init config failed:", err)
			return err
		}
	}

	conf := filepath.Join(baseSpace, ".kube/config")
	if cluster, err := k8s.NewCluster(conf, "default-system", b.Namespace); err != nil {
		return err
	} else {
		b.cluster = cluster
	}

	return nil
}

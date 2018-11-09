package builder

import (
	"fmt"
	// "k8s.io/apimachinery/pkg/api/resource"
	"encoding/json"
	"errors"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sort"
	"strconv"
	"strings"
)

const (
	DeployVersion = "tencentcloud-deploy-version"
	DeployGroup   = "tencentcloud-deploy-group"
	ServiceGate   = "tencentcloud-service-gate"

	PatchTypeMetaAndTemplate = "meta_template"
	PatchTypeMataAndSelector = "meta_selector"
	PatchTypeTemplate        = "template"
	PatchTypeAll             = "all"

	StrategyRecreate  = "recreate"
	StrategyBlueGreen = "blue-green"
	StrategyCanary    = "canary"
	StrategyOffline    = "offline"
)

// 初始化遗留部署
func (c *Cluster) TryInitDeployment(dm *extv1beta1.Deployment, deployGroup string, services []string) error {
	fmt.Printf("try TryInitDeploymentVersion with args: %s %s %s\n", dm.Name, deployGroup, services)

	labels := make(map[string]string)

	if deployGroup != "" && dm.ObjectMeta.Labels[DeployGroup] == "" {
		labels[DeployGroup] = deployGroup
	}

	if dm.ObjectMeta.Labels[DeployVersion] == "" {
		labels[DeployVersion] = "0"
	}

	if len(labels) > 0 {
		change, _ := json.Marshal(labels)
		err := c.AddLabelsToDeployment(dm, string(change), PatchTypeAll)
		if err != nil {
			return err
		}
	}

	// deployment service
	labels = make(map[string]string)

	for _, name := range services {
		gateKey := fmt.Sprintf("%s-%s", ServiceGate, name)
		labels[gateKey] = "true"
	}
	if len(labels) > 0 {
		change, _ := json.Marshal(labels)
		err := c.AddLabelsToDeployment(dm, string(change), PatchTypeMetaAndTemplate)
		if err != nil {
			return err
		}
	}

	// service
	for _, name := range services {
		service, err := c.GetService(name)
		if err != nil {
			return err
		}
		_, err = c.InitService(service)
		if err != nil {
			return err
		}
	}

	//if len() != "" && dm.ObjectMeta.Labels[enabledSvcName] != "true" {
	//	return c.EnableDeploymentTraffic(dm, enabledSvcName)
	//}
	return nil
}

// Getter

func (c *Cluster) GetDeployment(name string) (*extv1beta1.Deployment, error) {
	fmt.Printf("try GetDeployment with args: %s \n", name)
	return c.deploymentInterface.Get(name, metav1.GetOptions{})
}

func (c *Cluster) GetDeploymentByDeployGroup(dgName, target string) (dm *extv1beta1.Deployment, err error) {
	fmt.Printf("try GetDeploymentByDeployGroup  with args: %s %s\n", dgName, target)
	dmList, err := c.ListAllDeploymentsByDeployGroup(dgName)
	if err != nil {
		return
	}
	if len(dmList) == 0 {
		// return nil, errors.New("no deployment")
		return c.GetDeployment(dgName) // 临时方案, 对tke的服务初次部署时, 尝试查找和dgName同名的deployment
		// return nil, nil
	}
	if target == "newest" {
		dm = dmList[len(dmList)-1]
	} else if target == "previous" {
		if len(dmList) < 2 {
			// return nil, errors.New("no previous deployment")
			return nil, nil
		}
		dm = dmList[len(dmList)-2]
	} else if target == "oldest" {
		dm = dmList[0]
	} else {
		return nil, errors.New("暂时只支持 newest, previous, oldest")
	}

	return
}

// List Getter

func (c *Cluster) ListDeployments(selector string) (*extv1beta1.DeploymentList, error) {
	fmt.Printf("try ListDeployments with args: %s\n", selector)
	return c.deploymentInterface.List(metav1.ListOptions{LabelSelector: selector})
}

func (c *Cluster) ListAllDeploymentsByDeployGroup(deployGroup string) (dmList DeploymentList, err error) {
	fmt.Printf("try ListAllDeploymentsByDeployGroup with args: %s\n", deployGroup)
	dms, err := c.ListDeployments(fmt.Sprintf("%s=%s", DeployGroup, deployGroup))
	if err != nil {
		return
	}
	for i := range dms.Items {
		dmList = append(dmList, &dms.Items[i])
	}
	sort.Sort(dmList)

	return
}

// Attr Getter

func (c *Cluster) GetDeploymentVersion(dm *extv1beta1.Deployment) (int64, error) {
	fmt.Printf("try GetDeploymentVersion with args: %s\n", dm.Name)
	v, ok := dm.ObjectMeta.Labels[DeployVersion]

	if !ok {
		return 0, fmt.Errorf("deployment %s has no deploy version", dm.Name)
	}

	version, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}

	return version, nil
}

/*
func (c *Cluster) GetDeploymentServices(dm *extv1beta1.Deployment) []string {
	var services []string
	for k, v := range dm.ObjectMeta.Labels {
		if strings.HasPrefix(k, ServiceGate) && v == "true" {
			services = append(services, k[len(ServiceGate)+1:])
		}
	}
	return services
}
*/

// Create

func (c *Cluster) CreateDeployment(dm *extv1beta1.Deployment) error {
	fmt.Printf("try CreateDeployment with args: %s\n", dm.Name)
	dm.ObjectMeta.ResourceVersion = ""
	newDm, err := c.deploymentInterface.Create(dm)
	if err != nil {
		return err
	}
	*dm = *newDm
	return err
}

// Delete

func (c *Cluster) DeleteDeployment(dm *extv1beta1.Deployment) error {
	fmt.Printf("try DeleteDeployment with args: %s \n", dm.Name)
	policy := metav1.DeletePropagationBackground
	return c.deploymentInterface.Delete(dm.Name, &metav1.DeleteOptions{PropagationPolicy: &policy})
}

func (c *Cluster) EnableDeploymentAllTraffic(dm *extv1beta1.Deployment) error {
	fmt.Printf("try EnableDeploymentAllTraffic with args: %s\n", dm.Name)

	labels := make(map[string]string)
	for k, v := range dm.ObjectMeta.Labels {
		if strings.HasPrefix(k, ServiceGate) && v == "true" {
			labels[k] = "true"
		}
	}

	if len(labels) > 0 {
		change, err := json.Marshal(labels)
		if err != nil {
			return err
		}
		return c.AddLabelsToDeployment(dm, string(change), PatchTypeTemplate)
	}
	return nil
}

func (c *Cluster) DisableDeploymentAllTraffic(dm *extv1beta1.Deployment) error {
	fmt.Printf("try DisableDeploymentAllTraffic with args: %s\n", dm.Name)

	labels := make(map[string]string)
	for k, v := range dm.ObjectMeta.Labels {
		if strings.HasPrefix(k, ServiceGate) && v == "true" {
			labels[k] = "false"
		}
	}

	if len(labels) > 0 {
		change, err := json.Marshal(labels)
		if err != nil {
			return err
		}
		return c.AddLabelsToDeployment(dm, string(change), PatchTypeTemplate)
	}
	return nil
}

func (c *Cluster) PauseDeployment(dm *extv1beta1.Deployment) error {
	fmt.Printf("try PauseDeployment with args: %s\n", dm.Name)
	change := `{"spec": {"paused": true}}`
	newDm, err := c.deploymentInterface.Patch(dm.Name, types.StrategicMergePatchType, []byte(change))
	if err != nil {
		return err
	}
	*dm = *newDm
	return nil
}

func (c *Cluster) ResumeDeployment(dm *extv1beta1.Deployment) error {
	fmt.Printf("try ResumeDeployment with args: %s\n", dm.Name)
	change := `{"spec": {"paused": false}}`
	newDm, err := c.deploymentInterface.Patch(dm.Name, types.StrategicMergePatchType, []byte(change))
	if err != nil {
		return err
	}
	*dm = *newDm
	return nil
}

func (c *Cluster) CloneDeployment(info *DeployInfo) (*extv1beta1.Deployment, error) {
	//func (c *Cluster) CloneDeployment(templateDm *extv1beta1.Deployment, dgName, svcName, image string, replicas int32)(*extv1beta1.Deployment, error) {
	fmt.Printf("try CloneDeployment with args: %s\n", info.String())
	if info.DeployGroup == "" {
		return nil, errors.New("deploy group can not be empty")
	}
	if info.Name == "" {
		return nil, errors.New("deploy name can not be empty")
	}
	if info.Version == nil {
		return nil, errors.New("deploy version can not be empty")
	}

	newDm, err := c.GetDeployment(info.Template.Name)
	if err != nil {
		return nil, err
	}

	newDm.Name = info.Name
	newDm.Spec.Paused = false
	newDm.ObjectMeta.Labels[DeployGroup] = info.DeployGroup
	newDm.Spec.Template.ObjectMeta.Labels[DeployGroup] = info.DeployGroup
	newDm.Spec.Selector.MatchLabels[DeployGroup] = info.DeployGroup

	newDm.ObjectMeta.Labels[DeployVersion] = fmt.Sprintf("%d", *info.Version)
	newDm.Spec.Template.ObjectMeta.Labels[DeployVersion] = fmt.Sprintf("%d", *info.Version)
	newDm.Spec.Selector.MatchLabels[DeployVersion] = fmt.Sprintf("%d", *info.Version)

	if info.Image != "" {
		newDm.Spec.Template.Spec.Containers[0].Image = info.Image // TODO
	}

	for k, v := range newDm.ObjectMeta.Labels { // 清除之前的services
		if strings.HasPrefix(k, ServiceGate) && v == "true" {
			delete(newDm.ObjectMeta.Labels, k)
			delete(newDm.Spec.Template.ObjectMeta.Labels, k)
		}
	}

	gate := "true"
	if info.Strategy == StrategyOffline {
		gate = "false"
	}

	// for service lb
	for _, service := range info.Services {
		gateKey := fmt.Sprintf("%s-%s", ServiceGate, service)
		newDm.ObjectMeta.Labels[gateKey] = "true"
		newDm.Spec.Template.ObjectMeta.Labels[gateKey] = gate
	}

	if info.Replicas > 0 {
		newDm.Spec.Replicas = &info.Replicas
	}

	return newDm, nil
}

// TODO 处理默认策略
func (c *Cluster) Deploy(info *DeployInfo) (*extv1beta1.Deployment, error) {
	fmt.Printf("try Deploy with args: %s\n", info.String())

	if info.Template.ObjectMeta.Labels[DeployGroup] == "" {
		err := c.TryInitDeployment(info.Template, info.DeployGroup, info.Services)
		if err != nil {
			return nil, err
		}
	}

	dmList, err := c.ListAllDeploymentsByDeployGroup(info.DeployGroup)
	if err != nil {
		return nil, err
	}

	var newVersion int64
	var recentDm *extv1beta1.Deployment

	if len(dmList) > 0 {
		recentDm = dmList[len(dmList)-1]
		version, _ := c.GetDeploymentVersion(recentDm)
		newVersion = version + 1

		if version == 0 { // tke 历史部署
			info.Name = fmt.Sprintf("%s-%d", recentDm.Name, newVersion)
		} else {
			index := strings.LastIndex(recentDm.Name, "-")
			if index <= -1 {
				return nil, fmt.Errorf("can not parse name of deployment %s", recentDm.Name)
			}
			info.Name = fmt.Sprintf("%s-%d", recentDm.Name[0:index], newVersion)
		}
	} else {
		newVersion = 1 // 新部署版本从1开始
		info.Name = fmt.Sprintf("%s-%d", info.DeployGroup, newVersion)
	}
	info.Version = &newVersion

	newDm, err := c.CloneDeployment(info)
	if err != nil {
		return nil, err
	}

	err = c.CreateDeployment(newDm)
	if err != nil {
		return nil, err
	}

	if info.Strategy == StrategyRecreate {
		if recentDm != nil {
			return newDm, c.DeleteDeployment(recentDm)
		}
	} else if info.Strategy == StrategyBlueGreen {
		if recentDm != nil {
			err := c.DisableDeploymentAllTraffic(recentDm) // TODO 只关掉recent够吗
			return newDm, err
		}
	}

	//else if info.Strategy == StrategyCanary {
	//} else if info.Strategy == StrategyStatic {

	return newDm, nil
	//
	//
	//
	//
	//
	//
	//
	//
	//if len(dms) == 0 {
	//	return nil, fmt.Errorf("deployment list of service %s is empty", svcName)
	//}
	//
	//recentDm := dms[len(dms)-1]
	//version, _ := c.GetDeploymentVersion(recentDm)
	//newVersion := version + 1
	//
	//objectMeta := metav1.ObjectMeta{
	//	Labels: recentDm.ObjectMeta.Labels,
	//}
	//objectMeta.Labels[DeployVersion] = fmt.Sprintf("%d", newVersion)
	//if newVersion == 1 {
	//	objectMeta.Name = fmt.Sprintf("%s-%d", recentDm.Name, newVersion)
	//} else {
	//	index := strings.LastIndex(recentDm.Name, "-")
	//	if index <= -1 {
	//		return nil, fmt.Errorf("can not parse name of deployment %s", recentDm.Name)
	//	}
	//	objectMeta.Name = fmt.Sprintf("%s-%d", recentDm.Name[0:index], newVersion)
	//}
	//
	//// "deployment.kubernetes.io/revision"
	//spec := recentDm.Spec
	//spec.Paused = false
	//spec.Template.Spec.Containers[0].Image = image // TODO
	//// for service lb
	//enableKey := fmt.Sprintf("%s-enable", svcName)
	//spec.Template.ObjectMeta.Labels[enableKey] = "true"
	//// spec.Selector.MatchLabels[enableKey] = "true"
	//// 防止dm 竞争
	//spec.Template.ObjectMeta.Labels[DeployVersion] = fmt.Sprintf("%d", newVersion)
	//spec.Selector.MatchLabels[DeployVersion] = fmt.Sprintf("%d", newVersion)
	////spec.Template.ObjectMeta.Labels[objectMeta.Name] = "true"
	////spec.Selector.MatchLabels[objectMeta.Name] = "true"
	//
	//if replicas > 0 {
	//	spec.Replicas = &replicas
	//}
	//
	//newDm := &extv1beta1.Deployment {
	//	ObjectMeta: objectMeta,
	//	Spec: spec,
	//}
	//
	//err := c.CreateDeployment(newDm)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if strategy ==  "highlander" {
	//	recentDm, _ := c.GetDeployment(recentDm.Name) // need reload
	//	return newDm, c.DeleteDeployment(recentDm)
	//} else if strategy ==  "blue/green" {
	//	recentDm, _ := c.GetDeployment(recentDm.Name)
	//	err := c.DisableDeploymentTraffic(recentDm, svcName)
	//	return newDm, err
	//} else if strategy ==  "canary" {
	//} else {
	//}
	//
	//return newDm , nil

	//if err := c.CreateDeployment(dm, image); err != nil {
	//	return err
	//}
	//
	//
	//labels := dms.ObjectMeta.Labels
	//labels["version"] = "100"
	//
	//dm := &extv1beta1.Deployment {
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name: "newfox",
	//		Labels: labels,
	//	},
	//	// Spec: extv1beta1.DeploymentSpec{ }
	//	Spec: dms.Spec,
	//}
	//
	//return c.CreateDeployment(dm)
}

func (c *Cluster) ScaleDeployment(dm *extv1beta1.Deployment, scaleTo int32, autoDeletion bool) error {
	fmt.Printf("try ScaleDeployment %s \n", dm.Name)
	if scaleTo == 0 && autoDeletion {
		return c.DeleteDeployment(dm)
	}

	change := fmt.Sprintf(`{"spec": {"replicas": %d}}`, scaleTo)

	newDm, err := c.deploymentInterface.Patch(dm.Name, types.StrategicMergePatchType, []byte(change))
	if err != nil {
		return err
	}
	*dm = *newDm
	return nil
}

func (c *Cluster) RollbackDeployment(fromDm, toDm *extv1beta1.Deployment) error {
	fmt.Printf("try RollbackDeployment with args: %s %s\n", fromDm.Name, toDm.Name)

	if err := c.EnableDeploymentAllTraffic(toDm); err != nil {
		return err
	}
	if err := c.DisableDeploymentAllTraffic(fromDm); err != nil {
		return err
	}

	return nil
}

func (c *Cluster) AddLabelsToDeployment(dm *extv1beta1.Deployment, labels, patchType string) error {
	fmt.Printf("try AddLabelsToDeployment %s to %s\n", labels, dm.Name)
	c.PauseDeployment(dm)
	defer c.ResumeDeployment(dm)

	selector, _ := metav1.LabelSelectorAsSelector(dm.Spec.Selector)
	pods, _ := c.ListPods(selector.String())

	for i := range pods.Items {
		c.AddLabelsToPod(&pods.Items[i], labels)
	}

	podTemplateHash := pods.Items[0].Labels["pod-template-hash"]
	replicaSets, _ := c.ListReplicaSets(fmt.Sprintf("pod-template-hash=%s", podTemplateHash))

	for i := range replicaSets.Items {
		c.AddLabelsToReplicaSet(&replicaSets.Items[i], labels, false, patchType)
	}

	var change string
	if patchType == PatchTypeMetaAndTemplate { // 和service映射
		change = fmt.Sprintf(`{"metadata":{"labels":%s},"spec":{"template":{"metadata":{"labels": %s}}}}`, labels, labels)
	} else if patchType == PatchTypeTemplate { // 启停流量
		change = fmt.Sprintf(`{"spec":{"template":{"metadata":{"labels": %s}}}}`, labels)
	} else if patchType == PatchTypeMataAndSelector {
		change = fmt.Sprintf(`{"metadata":{"labels":%s},"spec":{"selector":{"matchLabels":%s}}}`, labels, labels)
	} else if patchType == PatchTypeAll { // 版本信息
		change = fmt.Sprintf(`{"metadata":{"labels":%s},"spec":{"selector":{"matchLabels":%s},"template":{"metadata":{"labels": %s}}}}`,
			labels, labels, labels)
	} else {
		return fmt.Errorf("wrong PatchType")
	}

	//change := fmt.Sprintf(`
	//	{
	//	  "metadata": {"labels": %s},
	//	  "spec": {
	//		 "selector": {"matchLabels": %s},
	//		 "template": {
	//		   "metadata": {"labels": %s}
	//		 }
	//	  }
	//	}`, labels, labels, labels)

	newDm, err := c.deploymentInterface.Patch(dm.Name, types.StrategicMergePatchType, []byte(change))
	if err != nil {
		return err
	}
	*dm = *newDm
	return nil
}

/*
func (c *Cluster) GetDeploymentByTarget(svc *apiv1.Service, target string) (dm *extv1beta1.Deployment, err error) {
	fmt.Printf("try GetDeploymentByTarget of service %s %s\n", svc.Name, target)
	dmList, err := c.ListAllDeploymentsByService(svc)
	if err != nil {
		return
	}
	if len(dmList) == 0 {
		return nil, errors.New("no deployment")
	}
	if target == "newest" {
		dm = dmList[len(dmList)-1]
	} else if target == "previous" {
		if len(dmList) < 2 {
			return nil, errors.New("no previous deployment")
		}
		dm = dmList[len(dmList)-2]
	} else if target == "oldest" {
		dm = dmList[0]
	} else {
		return nil, errors.New("暂时只支持 newest, previous, oldest")
	}

	return
}
*/
/*
func (c *Cluster) EnableDeploymentTraffic(dm *extv1beta1.Deployment, svcName string) error {
	fmt.Printf("try EnableDeploymentTraffic %s of service %s\n", dm.Name, svcName )
	labels := fmt.Sprintf(`{"%s-enable": "true"}`, svcName)
	return c.AddLabelsToDeployment(dm, labels, PatchTypeMetaAndTemplate)
	//c.PauseDeployment(dm)
	//defer c.ResumeDeployment(dm)
	//
	//selector, _ := metav1.LabelSelectorAsSelector(dm.Spec.Selector)
	//pods, _ := c.ListPods(selector.String())
	//
	//for i := range pods.Items {
	//	c.EnablePodTraffic(&pods.Items[i], svcName)
	//}
	//
	//podTemplateHash := pods.Items[0].Labels["pod-template-hash"]
	//replicaSets, _ := c.ListReplicaSets(fmt.Sprintf("pod-template-hash=%s", podTemplateHash))
	//
	//for i := range replicaSets.Items {
	//	c.EnableReplicaSetTraffic(&replicaSets.Items[i], svcName, false)
	//}
	//
	//change := fmt.Sprintf(`
	//	{
     //     "metadata": {"labels": {"%s-enable": "true"}},
	//	  "spec": {
	//		 "selector": {"matchLabels": {"%s-enable": "true"}},
	//		 "template": {
	//		   "metadata": {"labels": {"%s-enable": "true"}}
	//		 }
	//	  }
	//	}`, svcName, svcName, svcName)
	//
	//if newDm, err := c.deploymentInterface.Patch(dm.Name, types.StrategicMergePatchType, []byte(change)); err == nil {
	//	*dm = *newDm
	//} else {
	//	return err
	//}
	//
	//return nil
}
*/

/*
func (c *Cluster) DisableDeploymentTraffic(dm *extv1beta1.Deployment, services string) error {
	fmt.Printf("try DisableDeploymentTraffic %s of services %s\n", dm.Name, services)
	labels := make(map[string]string)

	for _, service := range strings.Split(services, ",") {
		gateKey := fmt.Sprintf("%s-%s", ServiceGate, service)
		if dm.ObjectMeta.Labels[gateKey] == "true" {
			labels[gateKey] = "false"
		}
	}

	if len(labels) > 0 {
		change, err := json.Marshal(labels)
		if err != nil {
			return err
		}
		return c.AddLabelsToDeployment(dm, string(change), PatchTypeTemplate)
	}
	return nil
	//c.PauseDeployment(dm)
	//defer c.ResumeDeployment(dm)
	//
	//selector, _ := metav1.LabelSelectorAsSelector(dm.Spec.Selector)
	//pods, _ := c.ListPods(selector.String())
	//
	//for i := range pods.Items {
	//	c.EnablePodTraffic(&pods.Items[i], services)
	//}
	//
	//podTemplateHash := pods.Items[0].Labels["pod-template-hash"]
	//replicaSets, _ := c.ListReplicaSets(fmt.Sprintf("pod-template-hash=%s", podTemplateHash))
	//
	//for i := range replicaSets.Items {
	//	c.EnableReplicaSetTraffic(&replicaSets.Items[i], services, false)
	//}
	//
	//change := fmt.Sprintf(`
	//	{
    //     "metadata": {"labels": {"%s-enable": "false"}},
	//	  "spec": {
	//		 "selector": {"matchLabels": {"%s-enable": "false"}},
	//		 "template": {
	//		   "metadata": {"labels": {"%s-enable": "false"}}
	//		 }
	//	  }
	//	}`, services, services, services)
	//
	//if newDm, err := c.deploymentInterface.Patch(dm.Name, types.StrategicMergePatchType, []byte(change)); err == nil {
	//	*dm = *newDm
	//} else {
	//	return err
	//}
	//
	//return nil
}
*/

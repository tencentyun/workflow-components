package k8s

import (
	"fmt"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// replicaSet
func (c *Cluster) GetReplicaSet(name string) (*extv1beta1.ReplicaSet, error) {
	return c.replicaSetInterface.Get(name, metav1.GetOptions{})
}

func (c *Cluster) ListReplicaSets(labelSelector string) (*extv1beta1.ReplicaSetList, error) {
	return c.replicaSetInterface.List(metav1.ListOptions{LabelSelector: labelSelector})
}

func (c *Cluster) CreateReplicaSet(rs *extv1beta1.ReplicaSet) (*extv1beta1.ReplicaSet, error) {
	return c.replicaSetInterface.Create(rs)
}

/*
func (c *Cluster) EnableReplicaSetTraffic(rs *extv1beta1.ReplicaSet, serviceName string, effectSubResources bool)(*extv1beta1.ReplicaSet, error) {
	// TODO label 是否需要设置

	if effectSubResources {
		selector, _ := metav1.LabelSelectorAsSelector(rs.Spec.Selector)
		pods, _ := c.ListPods(selector.String())

		for i := range pods.Items {
			c.EnablePodTraffic(&pods.Items[i], serviceName)
		}
	}
	change := fmt.Sprintf(`
		{
          "metadata": {"labels": {"%s-enable": "true"}},
		  "spec": {
			 "template": {
			   "metadata": {"labels": {"%s-enable": "true"}}
			 }
		  }
		}`, serviceName, serviceName)
	return c.replicaSetInterface.Patch(rs.Name, types.StrategicMergePatchType, []byte(change))
}

func (c *Cluster) DisableReplicaSetTraffic(rs *extv1beta1.ReplicaSet, serviceName string, effectSubResources bool)(*extv1beta1.ReplicaSet, error) {
	if effectSubResources {
		selector, _ := metav1.LabelSelectorAsSelector(rs.Spec.Selector)
		pods, _ := c.ListPods(selector.String())

		for i := range pods.Items {
			c.DisablePodTraffic(&pods.Items[i], serviceName)
		}
	}
	change := fmt.Sprintf(`
		{
          "metadata": {"labels": {"%s-enable": "false"}},
		  "spec": {
			 "template": {
			   "metadata": {"labels": {"%s-enable": "false"}}
			 }
		  }
		}`, serviceName, serviceName)
	return c.replicaSetInterface.Patch(rs.Name, types.StrategicMergePatchType, []byte(change))
}
*/

func (c *Cluster) AddLabelsToReplicaSet(rs *extv1beta1.ReplicaSet, labels string, effectSubResources bool, patchType string) error {
	if effectSubResources {
		selector, _ := metav1.LabelSelectorAsSelector(rs.Spec.Selector)
		pods, _ := c.ListPods(selector.String())

		for i := range pods.Items {
			c.AddLabelsToPod(&pods.Items[i], labels)
		}
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

	newRs, err := c.replicaSetInterface.Patch(rs.Name, types.StrategicMergePatchType, []byte(change))
	if err != nil {
		return err
	}
	*rs = *newRs
	return nil
}

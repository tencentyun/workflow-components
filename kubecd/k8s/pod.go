package k8s

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (c *Cluster) GetPod(podName string) (*apiv1.Pod, error) {
	return c.podInterface.Get(podName, metav1.GetOptions{})
}

func (c *Cluster) ListPods(selector string) (*apiv1.PodList, error) {
	return c.podInterface.List(metav1.ListOptions{LabelSelector: selector})
}

func (c *Cluster) EnablePodTraffic(pod *apiv1.Pod, serviceName string) error {
	labels := fmt.Sprintf(`{"%s-enable": "true"}`, serviceName)
	return c.AddLabelsToPod(pod, labels)
}

func (c *Cluster) DisablePodTraffic(pod *apiv1.Pod, serviceName string) error {
	labels := fmt.Sprintf(`{"%s-enable": "false"}`, serviceName)
	return c.AddLabelsToPod(pod, labels)
}

func (c *Cluster) AddLabelsToPod(pod *apiv1.Pod, labels string) error {
	change := fmt.Sprintf(`{"metadata": {"labels": %s}}`, labels)

	newPod, err := c.podInterface.Patch(pod.Name, types.StrategicMergePatchType, []byte(change))
	if err != nil {
		return err
	}
	*pod = *newPod
	return nil
}

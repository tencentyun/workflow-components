package builder

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	"k8s.io/client-go/tools/clientcmd"
)

var clusters = map[string]*Cluster{}
var defaultContext string

type Cluster struct {
	namespace           string
	podInterface        v1.PodInterface
	deploymentInterface v1beta1.DeploymentInterface
	replicaSetInterface v1beta1.ReplicaSetInterface
	serviceInterface    v1.ServiceInterface
	secretInterface     v1.SecretInterface
}

func NewCluster(configPath, context, ns string) (*Cluster, error) {
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: configPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	c := Cluster{
		podInterface:        clientSet.CoreV1().Pods(ns),
		deploymentInterface: clientSet.ExtensionsV1beta1().Deployments(ns),
		replicaSetInterface: clientSet.ExtensionsV1beta1().ReplicaSets(ns),
		serviceInterface:    clientSet.CoreV1().Services(ns),
		secretInterface:     clientSet.CoreV1().Secrets(ns),
	}

	return &c, nil
}

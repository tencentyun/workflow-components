package builder

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// service
func (c *Cluster) GetService(name string) (*apiv1.Service, error) {
	return c.serviceInterface.Get(name, metav1.GetOptions{})
}

func (c *Cluster) InitService(svc *apiv1.Service) (*apiv1.Service, error) {
	gateKey := fmt.Sprintf("%s-%s", ServiceGate, svc.Name)

	if _, ok := svc.Spec.Selector[gateKey]; ok {
		return svc, nil
	}

	change := fmt.Sprintf(`
		{
		  "spec": {
			 "selector": { "%s": "true"}
			 }
		  }
		}`, gateKey)

	return c.serviceInterface.Patch(svc.Name, types.StrategicMergePatchType, []byte(change))
}

// 版本回滚和服务没关系
//func (c *Cluster) RollbackService(svc *apiv1.Service, dm *extv1beta1.Deployment) error {
//	fmt.Printf("try RollbackService %s to deployment %s\n", svc.Name, dm.Name)
//	dms, err := c.ListEnabledDeploymentsByService(svc)
//
//	if err != nil {
//		return err
//	}
//
//	if err = c.EnableDeploymentTraffic(dm, svc.Name); err != nil {
//		return err
//	}
//
//	for _, enabledDm := range dms {
//		if enabledDm.Name == dm.Name {
//			continue
//		}
//		if err := c.DisableDeploymentTraffic(enabledDm, svc.Name); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

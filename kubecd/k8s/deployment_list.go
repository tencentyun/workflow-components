package k8s

import (
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	"strconv"
)

type DeploymentList []*extv1beta1.Deployment

func (dl DeploymentList) Swap(i, j int) {
	dl[i], dl[j] = dl[j], dl[i]
}

func (dl DeploymentList) Len() int {
	return len(dl)
}

func (dl DeploymentList) Less(i, j int) bool {
	di, _ := strconv.ParseInt(dl[i].ObjectMeta.Labels[DeployVersion], 10, 64)
	dj, _ := strconv.ParseInt(dl[j].ObjectMeta.Labels[DeployVersion], 10, 64)
	return di < dj
}

## 组件名称：CD

### CD

### 组件参数
#### 入参

#### 出参

### 源码地址

### 构建:

```
docker build -t hub.tencentyun.com/tencenthub/kubecd_deploy -f dockerfiles/deploy.Dockerfile .
docker build -t hub.tencentyun.com/tencenthub/kubecd_scale -f dockerfiles/scale.Dockerfile .
docker build -t hub.tencentyun.com/tencenthub/kubecd_disable -f dockerfiles/disable.Dockerfile .
docker build -t hub.tencentyun.com/tencenthub/kubecd_enable -f dockerfiles/enable.Dockerfile .
docker build -t hub.tencentyun.com/tencenthub/kubecd_delete -f dockerfiles/delete.Dockerfile .
docker build -t hub.tencentyun.com/tencenthub/kubecd_shrink -f dockerfiles/shrink.Dockerfile .
docker build -t hub.tencentyun.com/tencenthub/kubecd_rollback -f dockerfiles/rollback.Dockerfile .
```

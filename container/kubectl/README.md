## 组件名称：Kubectl

### Kubectl:

执行 kubectl 命令

### 组件参数
#### 入参

- `USERNAME` 必填，kubernetes 用户名
- `PASSWORD` 必填，kubernetes 用户密码
- `CERTIFICATE` 必填, kubernetes 证书内容
- `SERVER` 必填, kubernetes 服务器地址
- `COMMAND` 必填, 需要执行的 kubectl 命令及参数

腾讯云 kubernetes 容器集群用户, 可通过以下方式获取集群账号密码以及证书信息:

1) 登录 [容器服务控制台 > 集群](https://console.cloud.tencent.com/ccs)，单击需要连接的集群 ID/名称，查看集群详情。

2) 在集群信息页，单击【显示凭证】，查看用户名、密码和证书信息。

3) 复制或下载证书文件到本地

#### 出参

无

### 源码地址

[Kubectl](https://github.com/tencentyun/workflow-components/tree/master/container/kubectl)

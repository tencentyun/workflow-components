## Helm CMD

### Helm CMD:

在预装 `helm`, `kubectl` 的环境里执行用户自定义的shell命令. 该组件还支持自定义多个环境变量, 这些环境变量在`CMD`中可以直接使用.

### 组件参数
#### 入参

- `USERNAME` 选填，kubernetes 用户名
- `PASSWORD` 选填 用户密码
- `TOKEN` 选填，kubernetes token （用户名密码和token必填一种)
- `CERTIFICATE` 必填, kubernetes 证书内容
- `SERVER` 必填, kubernetes 服务器地址
- `CMD` 必填, 用户自定义shell命令, 支持多行, 使用`/bin/sh -c`执行
- 其他自定义入参: 可以在`CMD`中通过环境变量读取
  

 

腾讯云 kubernetes 容器集群用户, 可通过以下方式获取集群账号密码以及证书信息:

1) 登录 [容器服务控制台 > 集群](https://console.cloud.tencent.com/ccs)，单击需要连接的集群 ID/名称，查看集群详情。

2) 在集群信息页，单击【显示凭证】，查看用户名、密码和证书信息。

3) 复制或下载证书文件到本地

#### 出参
无

### 源码地址

Helm CMD: <https://github.com/tencentyun/workflow-components/tree/master/container/helm_cmd>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/helm_cmd .`

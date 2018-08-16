## 组件名称：Update TKE Service

### Update TKE Service:

通过调用TKE API [修改服务镜像](https://cloud.tencent.com/document/api/457/9630), 更新TKE服务的镜像, 可以用于镜像更新后触发服务更新

### 组件参数
#### 入参

- `TENCENTCLOUD_SECRET_ID` 必填，在[云API密钥管理](https://console.cloud.tencent.com/cam/capi)上申请的标识身份的SecretId，一个SecretId对应唯一的SecretKey
- `TENCENTCLOUD_SECRET_KEY` 必填，SecretId 对应的唯一SecretKey
- `REGION` 必填, 区域参数，用来标识希望操作哪个区域的实例
- `CLUSTER_ID` 必填, 服务所在的TKE 集群ID
- `SERVICE_NAME` 必填, TKE 服务名
- `IMAGE` 可选, 新镜像，如果服务中一个实例下只有一个容器可以传此参数(image和containers二者必填一个)
- `CONTAINERS` 可选, 新镜像，如果服务中一个实例下有多个容器需要传入此参数，需要一个合法的json字符串, 格式例如`{"containerName1": "image1", "containerName2": "image2"}`
- `NAMESPACE` 可选, kubernetes 服务命名空间, 默认为`default`

#### 出参
无

### 源码地址

Update TKE Service: <https://github.com/tencentyun/workflow-components/tree/master/container/update_tke_service>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/update_tke_service .`

## 组件名称：Copy Image

该组件可以实现将镜像从一个仓库复制到另一个仓库, 仓库可以属于不同的用户或组织.

### 组件参数

#### 入参


- `HUB_USER` 必填，来源镜像仓库用户名, `docker login` 镜像仓库的用户名
- `HUB_TOKEN` 必填，来源镜像仓库用户token, `docker login` 镜像仓库的用户密码
- `_WORKFLOW_FLAG_HUB_TOKEN` 可选, 工作流系统标志位, 此标志位如果为`true`, 将自动注入工作流用户名和密码, 因此可以省略`HUB_USER`和`HUB_TOKEN`
- `IMAGE` 必填, 来源镜像, 如`hub.cloud.tencent.com/fox/from_my_awesome_image`
- `TO_IMAGE` 必填, 复制的目标镜像, 如`hub.cloud.tencent.com/fox/to_my_awesome_image`
- `TO_HUB_USER` 必填，复制的镜像仓库用户名, `docker login` 镜像仓库的用户名
- `TO_HUB_TOKEN` 必填，复制的镜像仓库用户token, `docker login` 镜像仓库的用户密码
- `_WORKFLOW_FLAG_DIND` 必填, 工作流系统标志位, 该组件需要docker client运行环境, 将此标志位置为`true`将提供相关docker 运行环境


#### 出参

- `IMAGE_ID` 复制后新镜像生成的Image ID, 如`sha256:9aa1f5d00769e83ed75c0f7347990246eb71aa56403fb1769bc87988d9b1cb8f`
- `IMAGE_TAG_DIGEST` 复制后新镜像生成的Digest, 如`sha256:cc6521a0a1423161def9ac34437c45cbfb180581b8d0a006dba2c4be939f2f76`

### 源码地址

[Copy Image](https://github.com/tencentyun/workflow-components/tree/master/container/copy_image)

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/copy_image .`

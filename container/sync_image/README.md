## 组件名称：Sync Image

该组件可以实现将镜像从一个仓库同步到另一个仓库, 支持共有仓库和私有仓库.

### 组件参数

#### 入参


- `FROM_HUB_USER` 选填，来源私有镜像仓库用户名, 公共仓库则可为空, `docker login` 镜像仓库的用户名，
- `FROM_HUB_TOKEN` 选填，来源私有镜像仓库用户token, 公共仓库则可为空, `docker login` 镜像仓库的用户密码
- `_WORKFLOW_FLAG_HUB_TOKEN` 可选, 工作流系统标志位, 此标志位如果为`true`, 将自动注入工作流用户名和密码, 因此可以省略`_WORKFLOW_HUB_USER`和`_WORKFLOW_HUB_TOKEN`
- `FROM_IMAGE` 必填, 来源镜像, 如`hub.cloud.tencent.com/fox/from_my_awesome_image`
- `TO_IMAGE` 必填, 复制的目标镜像, 如`hub.cloud.tencent.com/fox/to_my_awesome_image`
- `TO_HUB_USER` 必填，复制的镜像仓库用户名, `docker login` 镜像仓库的用户名
- `TO_HUB_TOKEN` 必填，复制的镜像仓库用户token, `docker login` 镜像仓库的用户密码

#### 出参

- `IMAGE_ID` 复制后新镜像生成的Image ID, 如`sha256:9aa1f5d00769e83ed75c0f7347990246eb71aa56403fb1769bc87988d9b1cb8f`
- `IMAGE_TAG_DIGEST` 复制后新镜像生成的Digest, 如`sha256:cc6521a0a1423161def9ac34437c45cbfb180581b8d0a006dba2c4be939f2f76`

### 源码地址

Sync Image: <https://github.com/tencentyun/workflow-components/tree/master/container/sync_image>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/sync_image .`

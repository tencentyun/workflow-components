## 组件名称：Docker Builder

### Docker：
[Docker]根据Dockerfile生成Docker Image的组件, 并推送到image相关的镜像仓库中, 支持推送到其他平台仓库

### 组件参数

#### 入参

- `HUB_USER` 必填，镜像仓库用户名, `docker login` 镜像仓库的用户名
- `HUB_TOKEN` 必填，镜像仓库用户token, `docker login` 镜像仓库的用户密码
- `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: `${_WORKFLOW_GIT_CLONE_URL}`
- `GIT_REF` 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master 
- `GIT_TYPE` 非必填, 标识`GIT_REF`的类型: `branch`, `tag` 或者 `commit`
- `IMAGE` 必填, 目标镜像, 如`hub.cloud.tencent.com/fox/my_awesome_image`
- `IMAGE_ID` 非必填, 目标镜像Tag, 默认值`latest`
- `EXTRA_IMAGE_TAG` 非必填, 新增目标镜像Tag
- `IMAGE_TAG_FORMAT` 非必填, 新增镜像Tag的命名格式, 可以使用git 分支名/tag名/commit ID/当前时间等作为命名组成, 如`latest-$branch-$commit-$time`
- `BUILD_WORKDIR` 非必填, 默认".", 工作路径
- `DOCKERFILE_PATH` 非必填，默认"Dockerfile"，Dockerfile路径
- `NO_CACHE` 非必填，默认值是false, 如果为true 将开启`--no-cache`禁用 docker 构建缓存标志
- `BUILD_ARGS` 非必填，传递给`--build-arg`的构建参数, 必须是一个有效的json字符串, 如`{"HTTP_PROXY":"http://10.20.30.2:1234","TIMEOUT":"10"}`, 将生成构建参数`--build-arg HTTP_PROXY=http://10.20.30.2:1234 --build-arg TIMEOUT=10`


#### 出参

- `IMAGE` 构建的镜像地址, 不包含tag名称, 例如: `hub.cloud.tencent.com/fox/my_awesome_image`
- `IMAGE_TAG` 构建镜像生成的Tag, 例如: `latest`
- `IMAGE_ID` 构建镜像生成的Image ID, 如`sha256:9aa1f5d00769e83ed75c0f7347990246eb71aa56403fb1769bc87988d9b1cb8f`
- `IMAGE_TAG_DIGEST` 构建镜像生成的Digest, 如`sha256:cc6521a0a1423161def9ac34437c45cbfb180581b8d0a006dba2c4be939f2f76`

### 源码地址

Docker Builder: <https://github.com/tencentyun/workflow-components/tree/master/container/docker_builder>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/docker_builder .`

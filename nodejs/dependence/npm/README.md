## 组件名称：Nodejs NPM

使用 Node.js 依赖管理工具 NPM 安装项目依赖, 可选地将依赖上传到指定的二进制仓库.

### 组件参数
#### 入参
- `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: `${_WORKFLOW_GIT_CLONE_URL}`
- `GIT_REF` 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master
- `NPM_PARAMS` 非必填，如 `--save-dev`
- `HUB_REPO` 非必填，构建产物目标二进制仓库, 组件在构建完成后将尝试自动上传构建产物到此仓库, 如果此值为空将不会自动上传
- `HUB_USER` 非必填，上传构建产物对应仓库的用户名, 如果想使用当前主账号, 可以直接设置`_WORKFLOW_FLAG_HUB_TOKEN: true`, 执行引擎将自动注入当前用户名和token
- `HUB_TOKEN` 非必填，上传构建产物对应仓库的用户密码或者token, 同上, 如果设置了`_WORKFLOW_FLAG_HUB_TOKEN: true`, 此入参可以省略
- `ARTIFACT_PATH` 非必填，上传构建产物对应的仓库目录, 默认是仓库根目录
- `ARTIFACT_TAG` 非必填，上传构建产物对应的tag, 默认是latest

#### 出参
无

### Tag列表及其Dockerfile链接

* 10.9, latest: [Dockerfile](https://github.com/tencentyun/workflow-components/blob/512530040cec72325b7cb42e862b79fe60898f56/nodejs/dependence/npm/Dockerfile)

### 源码地址

Nodejs NPM: <https://github.com/tencentyun/workflow-components/tree/master/nodejs/dependence/npm>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/nodejs_npm:latest .`

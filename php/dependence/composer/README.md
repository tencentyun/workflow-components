## 组件名称：PHP Composer

### Composer：
使用PHP 依赖管理工具Composer, 安装项目依赖并上传到指定的二进制仓库.

### 组件参数
#### 入参
- `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: `${_WORKFLOW_GIT_CLONE_URL}`
- `GIT_REF` 非必填，源代码目标提交号或者分支名, 默认为master
- `HUB_REPO` 非必填，构建产物目标二进制仓库, 组件在构建完成后将尝试自动上传构建产物到此仓库, 如果此值为空将不会自动上传
- `HUB_USER` 非必填，上传构建产物对应仓库的用户名
- `HUB_TOKEN` 非必填，上传构建产物对应仓库的用户密码或者token
- `ARTIFACT_PATH` 非必填，上传构建产物对应的仓库目录, 默认是仓库根目录
- `ARTIFACT_TAG` 非必填，上传构建产物对应的tag, 默认是latest

#### 出参
- `ARTIFACT_URL` 成功上传的构建产物url

### 源码地址

[PHP Composer](https://github.com/tencentyun/workflow-components/tree/master/php/dependence/composer)

## 组件名称：PHP Composer

### Composer：
使用PHP 依赖管理工具Composer, 安装项目依赖并上传到指定的二进制仓库.

### 组件参数
#### 入参
- `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: `${_WORKFLOW_GIT_CLONE_URL}`
- `GIT_REF` 非必填，源代码目标提交号或者分支名, 默认为master
- `HUB_REPO` 非必填，构建产物目标二进制仓库, 组件在构建完成后将尝试自动上传构建产物到此仓库, 如果此值为空将不会自动上传
- `HUB_USER` 非必填，上传构建产物对应仓库的用户名, 如果想使用当前主账号, 可以直接设置`_WORKFLOW_FLAG_HUB_TOKEN: true`, 执行引擎将自动注入当前用户名和token
- `HUB_TOKEN` 非必填，上传构建产物对应仓库的用户密码或者token, 同上, 如果设置了`_WORKFLOW_FLAG_HUB_TOKEN: true`, 此入参可以省略
- `ARTIFACT_PATH` 非必填，上传构建产物对应的仓库目录, 默认是仓库根目录
- `ARTIFACT_TAG` 非必填，上传构建产物对应的tag, 默认是latest

#### 出参
- `ARTIFACT_URL` 成功上传的构建产物url

### Tag列表及其Dockerfile链接

* 1.7.2, latest: [Dockerfile](https://github.com/tencentyun/workflow-components/blob/c587a3a7ba3632ab7422d2a08efd8bc60c93f5d2/php/dependence/composer/Dockerfile)

### 源码地址

PHP Composer: <https://github.com/tencentyun/workflow-components/tree/master/php/dependence/composer>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/php_composer:latest .`

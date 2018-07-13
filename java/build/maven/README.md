## 组件名称：Java Maven Build

### mavan：
[maven]java maven的编译构建工具。

### 组件参数
#### 入参
- `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权
- `GIT_REF` 非必填，源代码目标提交号或者分支名, 默认为master
- `GOALS` 非必填，maven 构建目标, 默认是`package`
- `POM_PATH` 非必填，pom 文件相对路径, 默认`./pom.xml`
- `HUB_BIN_REPO` 非必填，构建产物目标二进制仓库, 组件在构建完成后将尝试自动上传构建产物到此仓库, 如果此值为空将不会自动上传
- `HUB_USER` 非必填，上传构建产物对应仓库的用户名
- `HUB_TOKEN` 非必填，上传构建产物对应仓库的用户密码或者token
- `BIN_PATH` 非必填，上传构建产物对应的仓库目录, 默认是仓库根目录
- `BIN_TAG` 非必填，上传构建产物对应的tag, 默认是latest

#### 出参
- `ARTIFACTS` 构建产物结果列表

### 源码地址

[Java Maven Build](https://github.com/tencentyun/workflow-components/tree/master/java/build/maven)

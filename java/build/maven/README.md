## 组件名称：Java Maven Build

### mavan：
[maven]java maven的编译构建工具。

### 组件参数
#### 入参
- `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权
- `GIT_REF` 非必填，源代码目标提交号或者分支名, 默认为master
- `GOALS` 非必填，maven 构建目标, 默认是`package`
- `POM_PATH` 非必填，pom 文件相对路径, 默认`./pom.xml`

#### 出参
- `ARTIFACTS` 构建产物结果列表

### 源码地址

[Java Maven Build](https://github.com/tencentyun/workflow-components/tree/master/java/build/maven)

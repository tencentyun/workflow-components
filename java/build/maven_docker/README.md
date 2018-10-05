## 组件名称：Java Build Maven

使用 maven 进行 java 项目构建, 可选地将构建产物上传到指定的二进制仓库.

### 组件参数
#### 入参
- `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: `${_WORKFLOW_GIT_CLONE_URL}`
- `GIT_REF` 非必填，源代码目标提交号或者分支名, 默认为master
- `M2_SETTINGS_XML` 非必填，`$user/.m2/setting.xml`配置文件内容，默认使用maven的全局配置
- `GOALS` 非必填，maven 构建目标, 默认是`package`
- `POM_PATH` 非必填，pom 文件相对路径, 默认`./pom.xml`
- `EXT_COMMAND` 非必填，GOALS之外的命令, 默认不执行

### Tag列表及其Dockerfile链接

* 3.5-jdk-8, latest: [Dockerfile](https://github.com/tencentyun/workflow-components/blob/c2d0c1ceb447694a092599858203d29dd877e6bb/java/build/maven/Dockerfile)

### 源码地址

Java Maven Build: <https://github.com/tencentyun/workflow-components/tree/master/java/build/maven>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/java_build_maven:latest .`

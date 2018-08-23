## 组件名称: Java Jdepend Analysis

### Java Jdepend Analysis:
Java Jdepend代码分析工具, 用于根据项目生成格式良好的度量标准报告

### 组件参数
#### 入参
* `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}
* `GIT_REF` 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master

#### 出参
无

### Tag列表及其Dockerfile链接
* 2.9.1, latest: [Dockerfile](https://github.com/tencentyun/workflow-components/blob/d9aceb59d41859bb833eb300355baef1420b9069/java/analysis/jdepend/Dockerfile)

### 源码地址

Java Junit Test: <https://github.com/tencentyun/workflow-components/tree/master/java/analysis/jdepend>

### 构建
`docker build -t hub.tencentyun.com/tencenthub/java_analysis_jdepend:latest .`
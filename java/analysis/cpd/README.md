## 组件名称：Java Cpd Analysis

### Java Cpd Analysis:
Java Cpd代码分析工具, 用于检查程序中是否存在重复代码

### 组件参数
#### 入参
* `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}
* `GIT_REF` 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master

#### 出参
无

### Tag列表及其Dockerfile链接
* 1.1, latest: [Dockerfile](https://github.com/tencentyun/workflow-components/blob/548c0eb6d83695c911011267ae5da7805e543631/java/analysis/cpd/Dockerfile)

### 源码地址

Java Cpd Analysis：<https://github.com/tencentyun/workflow-components/tree/master/java/analysis/cpd>

### 构建
`docker build -t hub.tencentyun.com/tencenthub/java_analysis_cpd:latest .`
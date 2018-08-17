## Golang Golint Analysis

Golang 语法规则和代码风格检测工具

## 组件参数
### 入参
* GIT_CLONE_URL 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}
* GIT_REF 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master
* LINT_PACKAGE 非必填, 待分析的代码包, 通过路径的形式给出, 默认检索所有的代码包
* LINT_PARAMS 非必填, golint参数，用于指明代码检测信息的输出格式

### 出参
无

##源码地址

[Golang Golang Analysis](https://github.com/tencentyun/workflow-components/tree/master/golang/analysis/golint)

## 构建
docker build -t hub.tencentyun.com/tencenthub/golang_analysis_golint:latest .
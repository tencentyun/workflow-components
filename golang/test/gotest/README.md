## Golang Gotest Analysis

Golang 单元测试工具

## 组件参数
### 入参
* GIT_CLONE_URL 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}
* GIT_REF 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master
* GTEST_PACKAGE 非必填, 待测试的代码包(多个代码包通过空格分割)或者文件(test文件和源码文件通过空格分割), 代码包通过路径的形式给出, 默认检索所有的代码包
* GTEST_PARAMS 非必填, 参数，用于指明代码测试信息的输出格式

### 出参
无

## 源码地址

[Golang Gotest Analysis](https://github.com/tencentyun/workflow-components/tree/master/golang/analysis/gotest)

## 构建
docker build -t hub.tencentyun.com/tencenthub/golang_analysis_gotest:latest .
## Java FindBugs Analysis

Java FindBugs代码分析工具

## 组件参数
### 入参
* GIT_CLONE_URL 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}
* GIT_REF 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master

### 出参
无

## 源码地址

Java Findbugs Analysis: <https://github.com/tencentyun/workflow-components/tree/master/java/analysis/findbugs>

## 构建
docker build -t hub.tencentyun.com/tencenthub/java_analysis_findbugs:latest .
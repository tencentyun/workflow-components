## 组件名称: Java PMD Analysis

### Java PMD Analysis:
源代码分析器PMD，用于找到常见的编程缺陷，如未使用的变量，空捕获块，不必要的对象创建等等

### 组件参数
#### 入参
* `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}
* `GIT_REF` 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master

#### 出参
无

## Tag列表及其Dockerfile链接
* 6.6.0, latest: [Dockerfile]()

### 源码地址

Java PMD Analysis: <https://github.com/tencentyun/workflow-components/tree/master/java/analysis/pmd>

### 构建
`docker build -t hub.tencentyun.com/tencenthub/java_analysis_pmd:latest .`
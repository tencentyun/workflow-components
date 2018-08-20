## 组件名称：Python Coala Analysis

使用 Coala 检查 Python 项目语法规则和代码风格

### 组件参数
#### 入参
- `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: `${_WORKFLOW_GIT_CLONE_URL}`
- `GIT_REF` 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master
- `FILES` 非必填，目标文件, 默认是项目下所有py文件
- `BEARS` 非必填，coala bears,  默认值: `PEP8Bear,PyUnusedCodeBear`

#### 出参
无

### 源码地址

Python Coala Analysis: <https://github.com/tencentyun/workflow-components/tree/master/python/analysis/coala>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/python_analysis_coala:latest .`
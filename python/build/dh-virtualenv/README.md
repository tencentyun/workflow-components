## 组件名称：Python Build Dh_virtualenv

### Python Build Dh_virtualenv:
使用 Dh_virtualenv 构建python代码

### 组件参数
#### 入参
- `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: `${_WORKFLOW_GIT_CLONE_URL}`
- `GIT_REF` 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master

#### 出参
无

### 源码地址

Python dh_virtualenv build: <https://github.com/tencentyun/workflow-components/tree/master/python/build/dh_virtualenv>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/python_build_dh_virtualenv:latest .`

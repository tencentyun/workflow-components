## 组件名称：Nodejs Babel Build

### Babel：
[Babel]用于编写下一代 JavaScript 的编译器。


### 组件参数
#### 入参
- `GIT_CLONE_URL` 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: `${_WORKFLOW_GIT_CLONE_URL}`
- `GIT_REF` 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master
- `BUILD_PATH` 必填，目标文件路径
- `BUILD_PARAMS` 非必填，如 `--out-dir public`
#### 出参
无

### 源码地址

[Nodejs Babel Build](https://github.com/tencentyun/workflow-components/tree/master/nodejs/build/babel)

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/nodejs_build_babel:latest .`

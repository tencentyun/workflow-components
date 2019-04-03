## Golang Build

Golang 构建工具

## 组件参数
### 入参
* GIT_CLONE_URL 必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}
* GIT_REF 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master
* BUILD_PACKAGE_NAME 必填，golang项目的包名，比如github.com/golang/dep
* PACKAGE_TARGET 非必填，main包所的目录, 相对路径, 比如./cmd/dep，默认为.
* BUILD_VENDOR_CMD 非必填, 下载必须的依赖包的vendor目录, 比如go get -v $(go list ./... | grep -v vendor)
* OUTPUT 非必填, 参数，构建产物的名字

### 出参
无

## 源码地址

[Golang Build](https://github.com/tencentyun/workflow-components/tree/master/golang/build)

## 构建
docker build -t hub.tencentyun.com/tencenthub/gobuild:latest . 
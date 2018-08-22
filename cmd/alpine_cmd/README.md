## 组件名称：Alpine CMD

### Alpine CMD:

在Alpine 环境里执行用户自定义的shell命令. 该组件还支持自定义多个环境变量, 这些环境变量在`CMD`中可以直接使用.

### 组件参数
#### 入参

- `CMD` 必填, 用户自定义shell命令, 支持多行, 使用`/bin/sh -c`执行
- 其他自定义入参: 可以在`CMD`中通过环境变量读取
  

#### 出参
无

### Tag列表及其Dockerfile链接

* 3.8, latest: [Dockerfile](https://github.com/tencentyun/workflow-components/blob/31a3144a145ae2ebab6efecafa26ccdb4d955463/cmd/alpine_cmd/Dockerfile)
* 3.7: [Dockerfile](https://github.com/tencentyun/workflow-components/blob/0ac19e1cac61bc9dfdd591a38bb5e41326f6f3c8/cmd/alpine_cmd/Dockerfile)

### 源码地址

Alpine CMD: <https://github.com/tencentyun/workflow-components/tree/master/cmd/alpine_cmd>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/alpine_cmd .`

## 组件名称：Mysql Client CMD

### Mysql Client CMD:

在预装 `mysql` 客户端的环境里执行用户自定义的shell命令. 该组件还支持自定义多个环境变量, 这些环境变量在`CMD`中可以直接使用.

### 组件参数
#### 入参

- `CMD` 必填, 用户自定义shell命令, 支持多行, 使用`/bin/sh -c`执行
- 其他自定义入参: 可以在`CMD`中通过环境变量读取
  

#### 出参
无

### 源码地址

Mysql Client CMD: <https://github.com/tencentyun/workflow-components/tree/master/cmd/mysql_client_cmd>

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/mysql_client_cmd .`

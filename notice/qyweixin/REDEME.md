## 组件名称：Notice Qyweixin

在Alpine环境里可以发送文本形式的信息到企业微信的个人用户、群组用户、标签用户。该组件还支持自定义多个环境变量，这些环境变量在CMD中可以直接使用。

### 组件参数

#### 入参

* CORP_ID 必填，企业微信上的企业ID
* APP_SECRET 必填，企业微信上的应用secret
* AGENT_ID 选填， 企业微信上的创建的应用ID
* USERS 必填，接受信息的个人用户，多个用户之间通过'|'分割
* PARTYS 选填，接受信息的群组用户，多个群组之间通过'|'分割
* TAGS 选填，接受信息的标签用户，多个标签之间通过'|'分割
* TEXT 选填，发送的信息内容

#### 出参

无

### 源码地址

[Notice Qyweixin](https://github.com/coderwangke/workflow-components)

### 构建

docker build -t hub.tencentyun.com/tencenthub/notice_qyweixin:latest .
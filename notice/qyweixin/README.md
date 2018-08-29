## 组件名称：Notice Qyweixin

在Alpine环境里可以发送文本形式的信息到企业微信的个人用户、群组用户、标签用户。该组件还支持自定义多个环境变量，这些环境变量在CMD中可以直接使用。

### 组件参数

#### 入参

* CORP_ID 必填，企业微信上的企业ID，[企业微信术语说明](https://work.weixin.qq.com/api/doc#10013)(企业ID、应用ID、应用secret)
* AGENT_ID 必填， 企业微信上的创建的应用ID
* APP_SECRET 必填，企业微信上的应用secret
* USERS 选填，接受信息的个人用户，多个用户之间通过'|'分割, 注意: USERS、PARTYS、TAGS不能同时为空
* PARTYS 选填，接受信息的群组用户，多个群组之间通过'|'分割
* TAGS 选填，接受信息的标签用户，多个标签之间通过'|'分割
* MESSAGE 选填，发送的信息内容
* _WORKFLOW_FLAG_TASK_DETAIL 非必填，如果为true, workflow会将当前工作流执行详情的json字符串注入到环境变量_WORKFLOW_TASK_DETAIL中

自定义消息(MESSAGE) 和发送工作流执行消息(_WORKFLOW_FLAG_TASK_DETAIL) 两者是二选一的关系, 自定义消息优先级更高, 当MESSAGE有值时,只发送用户自定义的文本消息.

#### 出参

无

### 源码地址

Notice Qyweixin: <https://github.com/tencentyun/workflow-components/tree/master/notice/qyweixin>

### 构建

docker build -t hub.tencentyun.com/tencenthub/notice_qyweixin:latest .
## 组件名称：Notice DingDing

[notice] 钉钉机器人通知组件

### 组件参数
#### 入参

- `WEBHOOK` 必填, 钉钉机器人Webhook地址, 如何配置钉钉自定义机器人并获得Webhook地址请参考: [钉钉自定义机器人](https://open-doc.dingtalk.com/docs/doc.htm?treeId=257&articleId=105735&docType=1)
- `_WORKFLOW_FLAG_TASK_DETAIL` 非必填，如果为true, workflow会将当前工作流执行详情的json字符串注入到环境变量`_WORKFLOW_TASK_DETAIL`中.
- `MESSAGE` 非必填，自定发送的文本消息
- `AT_MOBILES` 非必填，被@人的手机号, 仅对发送自定义文本消息有效(`MESSAGE`不为空)
- `IS_AT_ALL` 非必填，@所有人时:true, 否则为:false,  仅对发送自定义文本消息有效(`MESSAGE`不为空)

自定义消息(`MESSAGE`) 和发送工作流执行消息(`_WORKFLOW_FLAG_TASK_DETAIL`) 两者是二选一的关系, 自定义消息优先级更高, 当`MESSAGE` 有值时, 只发送用户自定义的文本消息.

#### 出参

无

### 源码地址

[Notice DingDing](https://github.com/tencentyun/workflow-components/tree/master/notice/dingding)

### 构建:

`docker build -t hub.tencentyun.com/tencenthub/notice_dingding:latest .`

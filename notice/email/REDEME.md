# 组件名称：Notice Email

在Alpine环境里可以发送文本形式的邮件。该组件还支持自定义多个环境变量，这些环境变量在CMD中可以直接使用。

## 组件参数

### 入参

* FROM_USER 必填，邮件的发送方，如xxyyzz@qq.com
* TO_USERS 必填，邮件的接收方, 如user_xx@qq.com,user_yy@163.com,多个收件人之间通过','分割
* SMTP_SERVER_PORT 必填，SMTP服务器和端口(smtp.example.com:123),根据发送方的邮箱服务提供商确定SMTP服务器，举例:QQ的SMTP是smtp.qq.com:465、163的SMTP是smtp.163.com:465、126的SMTP是smtp.126.com:25等
* SECRET 必填，SMTP服务器(邮件发送服务器)的授权码(如QQ和163邮箱)或者邮箱的登录密码(如126邮箱)
* SUBJECT 选填，邮件的主题
* TEXT 选填，邮件的内容
* _WORKFLOW_FLAG_TASK_DETAIL 非必填，如果为true, workflow会将当前工作流执行详情的json字符串注入到环境变量_WORKFLOW_TASK_DETAIL中

自定义内容(TEXT) 和发送工作流执行消息(_WORKFLOW_FLAG_TASK_DETAIL) 两者是二选一的关系, 自定义消息优先级更高, 当TEXT有值时,只发送用户自定义的文本消息.

### 出参

无

## 源码地址

[Notice Email](https://github.com/tencentyun/workflow-components/tree/master/notice/email)

## 构建

docker build -t hub.tencentyun.com/tencenthub/notice_email .
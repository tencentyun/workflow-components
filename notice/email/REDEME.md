# 组件名称：Notice Email

在Alpine环境里可以发送文本形式的邮件。该组件还支持自定义多个环境变量，这些环境变量在CMD中可以直接使用。

## 组件参数

### 入参

* FROM_USER 必填，邮件的发送方，如xxyyzz@qq.com
* TO_USERS 必填，邮件的接收方, 如user_xx@qq.com | user_yy@163.com，多个收件人之间通过'|'分割
* SMTP_SERVER 必填，SMTP服务器(邮件发送服务器)，根据发送方的邮箱服务提供商确定SMTP服务器，举例：QQ的SMTP是smtp.qq.com、163的SMTP是smtp.163.com、126的SMTP是smtp.126.com等
* SECRET 必填，SMTP服务器(邮件发送服务器)的授权码(如QQ和163邮箱)或者邮箱的登录密码(如126邮箱)
* SMTP_PORT 必填，SMTP服务器(邮件发送服务器)的端口
* SUBJECT 选填，邮件的主题
* TEXT 选填，邮件的内容

### 出参

无

## 源码地址

[Notice Email](https://github.com/coderwangke/workflow-components)

## 构建

docker build -t hub.tencentyun.com/tencenthub/email .
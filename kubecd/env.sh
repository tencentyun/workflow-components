export GOPATH=/Users/zhonghua/.gvm/pkgsets/go1.11/global:/Users/zhonghua/code/tx/goprojects
export USERNAME=admin
export SERVER=https://cls-3bqhvijc.ccs.tencent-cloud.com
export PASSWORD=lV1ZVMoZapqEDWhxt9gHPwqoUY0gkgW5
export CERTIFICATE=$(cat /Users/zhonghua/.kube/foxzhongtest.crt)
export NAMESPACE=zhong

#1 重建发布
# export ACTION=deploy
# export IMAGE=hub.tencentyun.com/foxzhong/web_service:green
# export STRATEGY=recreate
# export REPLICAS=2
# export DEPLOY_GROUP=fox-web-service
# export DEPLOYMENT_NAME=fox-web-service
# export SERVICES=fox-web-service

#2 蓝绿发布
# export ACTION=deploy
# export IMAGE=hub.tencentyun.com/foxzhong/web_service:blue
# export STRATEGY=blue-green
# export REPLICAS=2
# export DEPLOY_GROUP=fox-web-service
# export SERVICES=fox-web-service
# unset DEPLOYMENT_NAME

#3 灰度发布
# export ACTION=deploy
# export IMAGE=hub.tencentyun.com/foxzhong/web_service:red
# export STRATEGY=canary
# export REPLICAS=1
# export DEPLOY_GROUP=fox-web-service
# export SERVICES=fox-web-service
# export DEPLOY_TARGET=oldest
# unset DEPLOYMENT_NAME

#4扩容
# export ACTION=scale
# export DEPLOYMENT_NAME=fox-web-service-3
# export SCALE_TO=2


#4 离线发布
# export ACTION=deploy
# export IMAGE=hub.tencentyun.com/foxzhong/web_service:yellow
# export STRATEGY=offline
# export REPLICAS=2
# export DEPLOY_GROUP=fox-web-service
# export SERVICES=fox-web-service
# export DEPLOY_TARGET=previous
# unset DEPLOYMENT_NAME


#5收缩版本
# export ACTION=shrink
# export DEPLOY_GROUP=fox-web-service
# export SHRINK_TO=3

#6缩容
# export ACTION=scale
# export DEPLOY_GROUP=fox-web-service
# export DEPLOY_TARGET=oldest
# export SCALE_DOWN=1
# unset SCALE_TO
# unset SCALE_UP
# unset DEPLOYMENT_NAME

#7缩容(自动删除)
# export ACTION=scale
# export DEPLOY_GROUP=fox-web-service
# export DEPLOYMENT_NAME=fox-web-service-2
# export SCALE_DOWN=1
# unset SCALE_TO
# unset SCALE_UP
# export AUTO_DELETION=true

#11开启流量
# export ACTION=enable
# export DEPLOY_GROUP=fox-web-service
# export DEPLOY_TARGET=newest
# unset DEPLOYMENT_NAME

#7关闭流量
# export ACTION=disable
# export DEPLOY_GROUP=fox-web-service
# export DEPLOY_TARGET=oldest
# unset DEPLOYMENT_NAME

#9回滚
# export ACTION=rollback
# export FROM_DEPLOYMENT_NAME=fox-web-service-4
# export TO_DEPLOYMENT_NAME=fox-web-service-3

#10删除指定版本
export ACTION=delete
export DEPLOYMENT_NAME=fox-web-service-4

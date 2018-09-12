## 工作流组件: Qta Create
### Qta Create:
Qta创建工具

### 组件参数
#### 入参
* `NAME` 必填，测试计划名称
* `PRODUCT_PATH` 必填, 安卓包路径，传入方式：QCI构建后传入
* `TEST_REPO_URL` 必填, 测试代码路径，传入方式：QCI将用例库打包后传入给QTA
* `TESTCASENAME` 必填, 要执行的测试用例集：传入方式：用户自行传入，多个用例集以空格间隔

#### 出参
* `_WORKFLOW_TASK_PLAN_ID`  测试计划ID

### Tag列表及其Dockerfile链接
* 1.0, latest: [Dockerfile](https://github.com/tencentyun/workflow-components/blob/022eddcac14c3ff3b45f35a35f2020dc8d114855/service/qta_create/Dockerfile)

### 源码地址

Qta Create: <https://github.com/tencentyun/workflow-components/tree/master/service/qta_create>

### 构建
`docker build -t hub.tencentyun.com/tencenthub/qta_create:latest .`

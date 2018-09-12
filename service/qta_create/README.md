## 工作流组件:Qta Create
### Qta Create
Qta创建工具

### 组件参数
#### 入参
* `NAME` 必填，测试计划名称
* `PRODUCT_PATH` 必填, 安卓包路径，传入方式：QCI构建后传入
* `TEST_REPO_URL` 必填, 测试代码路径，传入方式：QCI将用例库打包后传入给QTA
* `TESTCASENAME` 必填, 要执行的测试用例集：传入方式：用户自行传入，多个用例集以空格间隔

#### 出参
无

### Tag列表及其Dockerfile链接
* 1.0, latest: Dockerfile

### 源码地址

Qta Create: <https://github.com/tencentyun/workflow-components/tree/master/qta/create>

### 构建
`docker build -t hub.tencentyun.com/tencenthub/qta_create:latest .`

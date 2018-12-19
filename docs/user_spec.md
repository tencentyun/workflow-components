## 创建工作流指引

TencentHub提供编排工作流的能力。

---

### 前提条件

登录账号为具有该仓库的admin（管理）或write（读写）权限时，才可以编辑/执行工作流。


### 工作流页面

1.登录 [Tencent Hub控制台](https://console.cloud.tencent.com/tencenthub/store/registry)。

2.单击左侧导航栏中的【仓库管理】，下拉选择对应的【空间】，选择仓库。

3.仓库页面点击【工作流】。

### 新建工作流

1.点击【新建】。

2.设置工作流基础设置。

![](https://main.qcloudimg.com/raw/4a2dfa21ebf6911da770799677a1fb56.png)

- **工作流名称**： 填写工作流名称。  
- **超时时间**：当工作流执行时间超过您设置的超市时间时，工作流将被取消执行。   
- **GitServer地址**：绑定代码源仓库地址，工作流将自动生成系统环境变量，在整个工作流的JOB中均可使用；并可设置代码更新时自动触发工作流。选择您需要构建容器镜像的仓库。  
- **触发方式**：复选模式，支持当 push 代码到某个分支或者新的 Tag 时，自动触发工作流。您也可以不选择自动触发，仅使用手动触发时指定的代码分支或Tag或Commitid。

注意：

如果选择了代码源地址，工作流将为您自动生成环境参数，在整个工作流的每个任务中都可以获取到这些环境参数，您也可以在工作流的编辑中使用这些环境参数。     

| 环境变量名称              | 含义                             | 生成场景                  | 内容说明                                                       |
|---------------------------|----------------------------------|---------------------------|----------------------------------------------------------------|
| `_WORKFLOW_BUILD_TYPE`    | 构建类型, 表示工作流被触发的方式 | 默认生成                  | manually: 手动触发; webhook: webhook触发; api: API触发         |
| `_WORKFLOW_GIT_CLONE_URL` | git 克隆地址                     | 当工作流和已授权的git关联 | 包含Basic Auth信息的git 克隆地址                               |
| `_WORKFLOW_GIT_REF`       | git 引用                         | 同上                      | 可以是git tag, git branch 或者 git commit                      |


3.新增Stage

设置该Stage下的Job执行顺序：串行或并行

4.新增Job

- **选择工作方式**： 现在工作流仅支持使用工作流镜像组件执行作为JOB工作方式。   
- **选择工作流镜像组件**：选择我的工作流组件或推荐工作流组件。
  - **i:** 我的工作流组件，制作方式参考[工作流组件规范](https://cloud.tencent.com/document/product/857/17227)
  - **i:** 推荐工作流组件，在列表中选取符合您需求的工作流，关键词标签有助于您的筛选，在工作流的描述页面可以看到该组件功能的具体描述。  
 ![](https://main.qcloudimg.com/raw/f8eb11692efbf4a5089a45f8f52281c9.png)

- **填写Job参数**：点击填写说明可以查看该工作流组件的JOB参数设置。
  - **i:** 填写value值。  
  ![](https://main.qcloudimg.com/raw/4e83680ab8017a086d1690899eefb57b.png)
  - **ii:** 选择value值映射关系，可以选择value为环境变量或其他JOB的输出值。  
  ![](https://main.qcloudimg.com/raw/ac8037b49711f7552a3c13d2fb47e5cf.png)

### 工作流模版
系统预置了一些工作流模版供参考。  
![](https://main.qcloudimg.com/raw/297f14438127c89886d5d7fd30cbb417.png)


### 环境变量设置
在工作流的每一个JOB中均可以取到全局环境变量。Key只能由字母、数字和下划线组成, 且不能为空。  
![](https://main.qcloudimg.com/raw/93d219c7c6940e07160ea731bea6389c.png)

### 组件特殊标志

为了简化用户操作, Workflow 将一些常见的环境需求抽象为特殊标志, 用户可以在job中将这些特殊标志作为环境变量传入, 以实现构建环境定义:


| Flag 名称                    | 可选值     | 默认值 | 含义说明                                                                                       |
|------------------------------|------------|--------|------------------------------------------------------------------------------------------------|
| `_WORKFLOW_FLAG_HUB_TOKEN`   | true/false | false  | 如果为`true`, 工作流将自动注入工作流用户名和密码到环境变量`HUB_USER`和`HUB_TOKEN`              |
| `_WORKFLOW_FLAG_DIND`        | true/false | false  | 如果为`true`, 将开启docker in docker 特性                                                      |
| `_WORKFLOW_FLAG_TASK_DETAIL` | true/false | false  | 如果为`true`, 工作流将自动注入工作流当前执行详情到环境变量`WORKFLOW_TASK_DETAIL中`             |
| `_WORKFLOW_FLAG_CACHE`       | true/false | false  | 如果为`true`, 工作流将开启job共享缓存特性, 工作流中的所有job可共享运行时目录`/workflow_cache/` |

### 共享缓存

同一条工作流中, Job之间可以共享工作缓存, 具体做法:

1. 对需要共享缓存的Job, 都增加入参: `_WORKFLOW_FLAG_CACHE` 为 `true`, 这样一来, 这些job在运行时可以共享同一份目录`/workflow_cache/`
2. 在这些Job使用的组件中, 上游job可以把产物/缓存等物件, 放到`/workflow_cache/`中, 下游job就可以直接从`/workflow_cache/`中读取.

## ego-layout

开发文档：http://www.mittacy.com/column/1633512445750

本项目是对 Gin 的二次封装，便于快速开发业务服务

> 1. 这是一个基于 Go 语言、Gin 框架的 Web 项目骨架，对常用的库进行封装，开发者可以更快速搭建 Web 服务，更关注属于自己的业务
> 2. 项目要求 Go 版本 >= 1.15
> 3. 拉取本项目骨架，在此基础上就可以快速开发自己的项目

### 1. 项目架构

<img src="README.assets/framework.png" alt="image-20210626172449172" style="zoom:50%;margin:0" />

- middleware：全局中间件，所有路由都会经过这些中间件

- router 层：定义路由，调用 api 层各个方法。每个路由也可以有自己的中间件。**建议分成两部分，高级权限的路由统一前缀管理，防止调用错路由**

- api层：包含一个或多个 service

    1. 调用 validator 层的请求结构体，解析请求参数，如果失败直接返回结果；
    2. 调用一个或多个 service 服务，获得返回结果；
    3. 调用 transform 对service返回的数据进行处理，然后响应给前端
    
- service 层：service 包含 data 或者调用其他 service
    1. 处理各种业务逻辑
    2. 调用 data 方法查询或存储数据、调用其他 service 完成操作
    
- data 层：包含db、cache、外部http调用服务等，涉及数据的查询和持久化都应该在该层实现

    > 不管是mysql、redis、mongo、es还是调用其他服务，这些操作其实都是数据的curd，对service层来说都是数据的查询与持久化，因此这些操作都应该写在data层

- model 层：定义数据库结构体、http远程调用响应结构体……

> model层与validator层的区别在于：
>
> + model层定义的是数据的原始结构体，即数据怎么存储那model的结构体就是一一对应的
> + validator层定义的结构体为请求与响应数据，即请求接口需要的数据结构、以及响应数据需要的数据结构


### 2. 项目结构

```shell
├── bootstrap					# 服务依赖初始化
│   ├── http.go					# http服务
│   ├── job.go					# 异步任务服务
│   ├── log.go
│   ├── task.go					# 定时任务服务
│   └── viper.go
├── apierr						# 业务码定义
│   └── code.go
├── cmd
│   └── start					
│       ├── http				# http
│       │   └── http.go
│       ├── job					# 异步任务
│       │   └── job.go
│       └── task				# 定时任务
│       │   └── task.go
│		└── start.go			# 服务启动程序
├── bin							# 可执行文件存储目录，不上传
├── config						# 配置设置
│   ├── async.go
│   ├── async_config
│   │   └── job.go
│   ├── mysql.go
│   ├── redis.go
│   └── task.go
├── middleware					# 中间件
│   ├── requestLog.go
│   └── requestTrace.go
├── router						# 路由
│   ├── admin.go
│   └── router.go
├── app
│   ├── api						# api控制器
│   │   └── user.go
│   ├── internal
│   │   ├── data
│   │   │   └── user.go
│   │   ├── model
│   │   │   └── user.go
│   │   ├── service
│   │   │   └── user.go
│   │   ├── transform			# 响应数据处理、封装
│   │   │   └── user.go
│   │   └── validator			# 数据请求与参数校验、响应结构体
│   │       └── userValidator
│   │           └── user.go
│   ├── job
│   │   ├── job_payload			# 异步任务数据定义
│   │   │   └── example.go
│   │   └── job_process			# 异步任务处理器
│   │       └── example.go
│   └── task
│       └── example.go
├── main.go
├── hook.go						# 服务钩子定义
├──.env							# 本地环境配置
├──.env.development				# 开发环境配置
├──.env.production				# 生产环境配置
```

## 快速开始

### 1. 环境准备

需要提前安装好对应的依赖环境以及工具：‌

- 安装go环境，version >= 1.16
- 安装Mysql（如果需要）
- 安装Redis（如果需要）

设置env

```shell
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

### 2. 修改配置信息

配置文件位于项目根目录 `.env.development`，修改对应的配置信息：

+ 服务信息
+ mysql信息
+ redis信息

### 3. 启动服务

#### 3.1 使用命令启动

```shell
$ cd myProjectName
$ go mod download
$ go build -o ./bin/server .

# 运行异步任务服务
$ ./bin/server start job -c=.env.development -e=development

# 运行定时任务
$ ./bin/server start task -c=.env.development -e=development

# 运行HTTP服务
$ ./bin/server start http -c=.env.development -e=development -p=8080

$ curl localhost:8080/api/user/ping
{"code":0,"data":"success","msg":"success","request_id":"r61f641e8bf370_pkL0LEODq4N2PyASnn"}
```

> Flags:
>   -c, --conf string   配置文件路径 (default ".env.development")
>   -e, --env string    运行环境 (default "development")
>
> + development 会将日志同步打印到控制台
> + test/production 不会将日志打印到控制台，正式环境应该设置为 production
>
>   -p, --port int      监听端口 (default 8080)

#### 3.2 Docker启动

```shell
$ docker build -t 镜像名 .

# 启动http服务
$ docker run --restart=on-failure --name 服务名1 -d -p 宿主端口:8080 镜像名 start http -c=.env.development -e=production -p=8080
# 简写
$ docker run --restart=on-failure --name 服务名1 -d -p 宿主端口:8080 镜像名

# 启动job异步服务
$ docker run --restart=on-failure --name 服务名2 -d -p 宿主端口:8080 镜像名 start job -c=.env.development -e=production

# 启动task定时服务
$ docker run --restart=on-failure --name 服务名2 -d -p 宿主端口:8080 镜像名 start task -c=.env.development -e=production
```

### 4. 代码生成

```shell
# 安装 ego 工具
$ go get -u github.com/mittacy/ego@latest

# 进入项目根目录
$ ego new project
$ cd project

# 创建 api、validator、transform、service、data、model 代码模板
$ ego tpl api article

# 创建 service、data、model 代码模板
$ ego tpl service article

# 创建 data、model 代码模板
$ ego tpl data article

# 创建定时任务 task 代码模板
$ ego tpl task notice

# 创建异步任务 job 代码模板
$ ego tpl job sendEmail
```

### 5. 插件

```shell
# 往项目注入git commit注释规范
$ ego plugin git -t=./.git -t=commitLint
```


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
- api层：可以包含一个或多个 service

    1. 调用 validator 层的请求结构体，解析请求参数，如果失败直接返回结果；
    2. 调用一个或多个 service 服务，获得返回结果；
    3. 调用 transform 对响应数据进行处理，然后响应给前端
- service 层：service 包含 data 以及其他 service，**注意不要循环造成调用服务**
    1. 处理各种逻辑
    2. 调用 data 方法存储或查询数据 或者 调用其他 service 完成操作
- data 层：包含db、cache、http远程调用服务等，涉及数据的查询和持久化都应该在该层实现
- model 层：定义数据库结构体、http远程调用响应结构体……


### 2. 项目结构

```shell
├── bootstrap					# 初始化封装
│   ├── http.go
│   ├── job.go
│   ├── log.go
│   ├── task.go
│   └── viper.go
├── apierr						# 服务错误码和错误定义
│   ├── code.go
│   └── err.go
├── cmd							# 服务命令
│   └── start					# start命令，启动http、job、task
│       ├── start.go
│       ├── http
│       │   └── http.go
│       ├── job
│       │   └── job.go
│       └── task
│           └── task.go					
├── app									# 服务
│   ├── api								# API控制器
│   │   └── user.go
│   ├── job								# 异步任务
│   │   └── exampleJob					
│   │       ├── exampleJobProcessor		# 任务处理器
│   │       │   └── processor.go
│   │       └── exampleJobTask			# 生成任务
│   │           └── task.go
│   └── task							# 定时任务
│       └── example.go
├── internal					# 内部服务
│   └── validator				# 数据请求、响应结构体定义以及参数校验
│   │   └── userValidator
│   │       └── user.go
│   ├── transform				# 响应数据处理、封装
│   │   └── user.go
│   ├── service					# 服务层，处理逻辑
│   │   └── user.go
│   ├── data					# 数据查询、存储层
│   │   └── user.go
│   └── model					# 定义与数据库的映射结构体
│       └── user.go          	
├── middleware					# 中间件
│   ├── requestLog.go			# 请求日志记录中间件
│   └── requestTrace.go			# 请求追踪中间件
└── router						# 路由
│   ├── admin.go
│   └── router.go
├── Makefile
├── main.go
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
$ go  build -o ./bin/server main.go

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
$ docker run --restart=on-failure --name 服务名1 -d -p 宿主端口:8080 镜像名 start http -c=.env.development -e=production -p 8080
# 简写
$ docker run --restart=on-failure --name 服务名1 -d -p 宿主端口:8080 镜像名

# 启动job异步服务
$ docker run --restart=on-failure --name 服务名2 -d -p 宿主端口:8080 镜像名 start job -c=.env.development -e=production

# 启动task定时服务
$ docker run --restart=on-failure --name 服务名2 -d -p 宿主端口:8080 镜像名 start task -c=.env.development -e=production
```



### 4. 生成业务框架代码

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


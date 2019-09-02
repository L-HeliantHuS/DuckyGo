# DuckyGo (基于gin+gorm搭建的快速高效稳定的web restful api)

项目基于MIT协议，任何人可以进行修改并发布，如果本项目你发现有任何BUG，欢迎提交PullRequest :fire:

## 目的 :star2:

本项目采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建Restful Web API

## 说明  :shipit:
响应序列化:
```text
任何一个json返回值，都是以serializer.Response进行序列化出来
最后在api(controll)层c.JSON返回的时候加.Result() 进行加TimeStamp
```

业务逻辑:
```text
api层不允许出现数据库操作, 所有的数据库操作都在service层进行. api层只进行大体的业务逻辑和序列化.
```

## 特色 :blue_heart:

本项目已经整合了许多开发API所必要的组件：

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的.
2. [GORM](http://gorm.io/docs/index.html): ORM工具。本项目需要配合Mysql使用.
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具.
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端.
5. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量.
6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件.
7. 橙卡大佬实现了国际化i18n的一些基本功能.
8. 本项目是使用Redis来保存用户Session登陆状态.
9. 使用Redis-list实现了内部消息队列,发送邮件可实现完全异步发送.

本项目已经预先实现了一些常用的代码方便参考和复用:

1. 创建了用户模型
2. 实现了```/api/v1/user/register```用户注册接口
3. 实现了```/api/v1/user/login```用户登录接口
4. 实现了```/api/v1/user/me```用户资料接口(需要登录后获取session)
5. 实现了```/api/v1/user/logout```用户登出接口(需要登录后获取session)
6. 实现了```/api/v1/user/changepassword```用户修改密码接口

本项目已经预先创建了一系列文件夹划分出下列模块:

1. api文件夹就是MVC框架的controller，负责协调各部件完成任务
2. model文件夹负责存储数据库模型和数据库操作相关的代码
3. service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象
5. cache负责redis, RabbitMQ缓存相关的代码
6. auth权限控制文件夹
7. util一些小工具, 目前有randomString、Logger、SendEmail
8. conf放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件
9. log放生成的日志文件，第一次使用需要双击运行一下bat文件生成log文件.

## LOG_LEVEL说明 :purple_heart:

第一次使用要先运行`log`文件夹下的`bat`批处理，用来生成记录log所需要的log文件.

```text

当设置LOG_LEVEL设置为ERROR
就只会显示 error panic

当设置LOG_LEVEL设置为WARNING
就只会显示 warning error panic

当设置LOG_LEVEL设置为INFO
就只会显示 info warning error panic

当设置LOG_LEVEL设置为DEBUG
则全部显示

```

## Godotenv :yellow_heart:

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_passwd@tcp(127.0.0.1:3306)/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接配置
RABBITMQ_DSN="amqp://mq_user:mq_passwd@localhost:5672/virtual_host"                             # RabbitMQ连接配置 默认没有开启
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW=""                 # Redis连接密码
REDIS_DB=""                 # Redis库从0到10，不填即为0
SESSION_SECRE=""            # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"            # 设置gin的运行模式，有 debug 和 release
LOG_LEVEL="ERROR"           # 设置为ERROR基本不会记录log 设置为DEBUG则会详细记录每次请求
```

Windows安装MySQL和Redis麻烦?:no_mouth: 你可以使用[Docker](https://hub.docker.com/)啊！:sunglasses:

- 快速起Redis: `docker run -di --name redis -p 6379:6379 redis` 
- 快速起MySQL: `docker run -di --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=yourpassword mysql` 

因为启动容器指定了--name, 后续可以使用`docker start|stop redis|mysql` 来进行开启或者关闭.

如需要使用navicat等工具管理MySQL，可能会出现报错等情况：:dizzy_face:
```shell
docker exec -it mysql /bin/bash    # 打开mysql bash交互
mysql -u root -p                   # 进入mysql交互
ALTER USER 'root'@'%' IDENTIFIED BY 'password' PASSWORD EXPIRE NEVER;            # 更改加密方式
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'yourpassword';   # 更改密码
FLUSH PRIVILEGES;                                                                # 刷新
```
    
## Go Mod :green_heart:

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init DuckyGo
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行 :heartpulse:

```shell
go run main.go
```

项目运行后启动在8000端口（可以修改，参考gin文档)   
本项目修改端口请查看`main.go`


## 编译 :exclamation:
```shell
go build main.go
```
如需交叉编译请看[这里](https://studygolang.com/articles/13760)

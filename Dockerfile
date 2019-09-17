FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /DuckyGo

WORKDIR /DuckyGo

# 编译成linux的可执行文件
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on GOARCH=amd64 go build -o duckygo

FROM alpine

# 配置环境变量 这里注意ip地址不可以填写127.0.0.1了~ 因为是在当前的容器里面 当前容器并没有 mysql 和 redis...
ENV MYSQL_DSN="db_user:db_passwd@tcp(172.168.70.171:8052)/db_name?charset=utf8&parseTime=True&loc=Local"
ENV REDIS_ADDR="172.168.70.171:8051"
ENV REDIS_PW=""
ENV REDIS_DB="0"
ENV SESSION_SECRE="4T8I9T48I094T389013h98r0PIK6Y5LUH5IJOWE"
ENV GIN_MODE="release"
ENV PORT=8000

# 创建翻译文件和日志文件目录  不要问为什么没有log日志文件夹 生产从不需要日志！！ (滑稽
RUN mkdir -p /www/conf/locales

WORKDIR /www

# 从上一个容器中将编译完毕的文件复制到这个容器中
COPY --from=build /DuckyGo/duckygo /usr/bin/duckygo

# 将翻译文件添加到本容器
ADD ./conf/locales /www/conf/locales

RUN chmod +x /usr/bin/duckygo

# 开放端口
EXPOSE 8000

# 执行
ENTRYPOINT ["duckygo"]
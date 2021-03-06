# msg-pusher

## msg-pusher是用Golang编写的轻量级高性能消息推送微服务

### 使用msg-pusher可以实现：

- 消息的实时推送
- 消息的定时推送

### 目前支持一下平台的推送：

- 阿里云短信
- 微信公众号
- 邮件服务

### 调用关系

![image](image/relationship.jpg)

### 核心模块

1. receiver

        用于接收需要推送的消息，并将消息存入缓存和消息队列
2. sender

        绑定消息队列，将得到数据推送到对应的服务商
3. corn

        定时将缓存中的数据持久化到数据库

### 服务集成

服务集成了 [prometheus](https://github.com/prometheus/prometheus)、[jaeger](https://github.com/jaegertracing/jaeger) ,你可以通过这些插件来观察msg-pusher的响应情况和性能状况

### 服务部署

msg-pusher依赖rabbit-mq、redis和mysql

- 二进制部署

```bash
cd msg-pusher
go build -o ./dist/sender ./cmd/sender/*.go
go build -o ./dist/receiver ./cmd/receiver/*.go

./dist/sender -f conf.yaml --log-path ./
./dist/receiver -f conf.yaml --log-path ./
```

- docker部署

```bash
docker run --name pusher -v $CONF_PATH:/app/msg-pusher/conf/conf.yaml -v $LOG_FILE_PATH:/app/msg-pusher/log -p 8990:8990 hiruok/msg-pusher:V0.1.0
```

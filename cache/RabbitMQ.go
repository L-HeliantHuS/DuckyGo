package cache

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"DuckyGo/util"
)

// 创建RabbitMQ连接单例
var Channel *amqp.Channel

// RabbitMQ 在中间件中初始化RabbitMQ连接
func InitRabbitMQ(amqpUrl string) {
	connection, err := amqp.Dial(amqpUrl)
	if err != nil {
		log.Fatal(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		util.Log().Panic(fmt.Sprint(err))
	}

	Channel = channel
}

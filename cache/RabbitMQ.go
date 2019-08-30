package cache

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"singo/util"
)

var Channel *amqp.Channel

func Init(amqpUrl string) {
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

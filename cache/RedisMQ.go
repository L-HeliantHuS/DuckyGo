package cache

import (
	"fmt"
	"log"
	"DuckyGo/util"
)

type RedisMessageQueue struct {
}

// 实例化全局RedisMQ单例 依赖Redis连接单例
var RedisMQ *RedisMessageQueue

// Publish 生产者
func (mq *RedisMessageQueue) Publish(queuename string, message string) error {
	err := RedisClient.LPush(queuename, message).Err()

	if err != nil {
		return err
	}

	return nil
}

// Custome 消费者
func (mq *RedisMessageQueue) Custome(queuename string, cb func(message string) error) error {
	go func() {
		for {
			message, err := RedisClient.BRPop(0, queuename).Result()

			if err != nil {
				util.Log().Error(fmt.Sprint(err))
			}

			err = cb(message[1])

			if err != nil {
				log.Printf("Execute Callback func Error: %s", err)
			}
		}
	}()

	return nil
}

// InitRedisMQ 实例化单例函数
func InitRedisMQ() {
	RedisMQ = new(RedisMessageQueue)
}

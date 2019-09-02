package main

import (
	"DuckyGo/conf"
	"DuckyGo/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()

	// 运行 起在8000端口
	r.Run(":8000")
}

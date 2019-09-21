package conf

import (
	"DuckyGo/cache"
	"DuckyGo/model"
	"DuckyGo/util"
	"github.com/gin-gonic/gin"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	if os.Getenv("GIN_MODE") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}

	if os.Getenv("RIM") == "use" {
		// 启动各种连接单例
		model.Database(os.Getenv("MYSQL_DSN"))
		cache.Redis()
		cache.InitRedisMQ()
		//cache.InitRabbitMQ(os.Getenv("RABBITMQ_DSN"))

		// 启动其他异步服务 (RedisMQ, RabbitMQ的应用

	}

}

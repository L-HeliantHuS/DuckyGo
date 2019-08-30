package conf

import (
	"github.com/gin-gonic/gin"
	"os"
	"singo/cache"
	"singo/model"
	"singo/util"

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

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
	cache.RabbitMQ(os.Getenv("RABBITMQ_DSN"))
}

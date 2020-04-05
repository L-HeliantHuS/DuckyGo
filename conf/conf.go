package conf

import (
	"DuckyGo/cache"
	"DuckyGo/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"time"
)

// 全局参数
var (
	SigningKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

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

	if gin.Mode() == gin.ReleaseMode {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println(`
			 ____             _           ____       
			|  _ \ _   _  ___| | ___   _ / ___| ___  
			| | | | | | |/ __| |/ / | | | |  _ / _ \ 
			| |_| | |_| | (__|   <| |_| | |_| | (_) |
			|____/ \__,_|\___|_|\_\\__, |\____|\___/ 
								   |___/             
			 服务器已经启动成功啦~  现在是Release模式~
				如果网站访问403, 请检查跨域配置哦.
		`)
		}()
	}

}

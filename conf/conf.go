package conf

import (
	"giligili/cache"
	"giligili/model"
	"giligili/tasks"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	// godotenv会解析 .env 文件，该文件是一个典型的INI格式的文件：
	// 在代码中调用 godotenv.Load() 即可解析并将相应的Key/Value对都放到环境变量中。
	// 例如可以通过 os.Getenv("PORT") 获取
	godotenv.Load()

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()

	// 启动定时任务
	tasks.CronJob()
}

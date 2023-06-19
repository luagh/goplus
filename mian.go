package main

import (
	"Goplus/bootstrap"
	btsConfig "Goplus/config"
	"Goplus/pkg/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}
func main() {
	//配置初始化 依赖命令行 --env参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
	// new 一个 Gin Engine 实例
	r := gin.New()
	//初始化数据库DB
	bootstrap.SetupDB()
	// 初始化路由绑定
	bootstrap.SetupRoute(r)

	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}

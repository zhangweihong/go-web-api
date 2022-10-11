package main

import (
	"gin-framework/basic/src/config"
	"gin-framework/basic/src/middleware"
	"gin-framework/basic/src/setting"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//跨域设置
	r.Use(middleware.Cors())
	//初始化设置
	setting.Init(r)
	//启动服务
	r.Run(config.ServerPort)
}

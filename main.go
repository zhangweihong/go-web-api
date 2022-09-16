package main

import (
	"gin-framework/basic/src/config"
	db "gin-framework/basic/src/database"
	"gin-framework/basic/src/middleware"
	"gin-framework/basic/src/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerToFile())

	//链接数据库
	db.ConnectDb()
	//设置路由
	router.SetRouter(r)
	//启动服务
	r.Run(config.ServerPort)

}

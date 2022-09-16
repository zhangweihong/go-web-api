package main

import (
	"gin-framework/basic/src/config"
	db "gin-framework/basic/src/database"
	"gin-framework/basic/src/router"
	"gin-framework/basic/src/setting"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//初始化设置
	setting.Init(r)
	//链接数据库
	db.ConnectDb()
	//设置路由
	router.SetRouter(r)
	//启动服务
	r.Run(config.ServerPort)
}

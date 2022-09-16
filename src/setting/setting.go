package setting

import (
	db "gin-framework/basic/src/database"
	"gin-framework/basic/src/middleware"
	"gin-framework/basic/src/router"
	"gin-framework/basic/src/schedule"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	db.ConnectDb()                   //链接数据库
	router.SetRouter(r)              //设置路由
	r.Use(middleware.LoggerToFile()) //开启日志中间件
	schedule.Start()                 //开始定时任务
}

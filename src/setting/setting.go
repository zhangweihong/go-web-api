package setting

import (
	"gin-framework/basic/src/middleware"
	"gin-framework/basic/src/schedule"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Use(middleware.LoggerToFile()) //开启日志中间件
	schedule.New()                   //开始定时任务
}

package router

import (
	"gin-framework/basic/src/middleware"

	"github.com/gin-gonic/gin"
)

//设置路由接口
func SetRouter(r *gin.Engine) *gin.Engine {
	r.Static("/static", "./public")    //设置静态资源
	r.Use(middleware.SetRouterGuard()) //设置路由监测
	UserRouter(r)                      //设置对应的User相关的路由设置
	return r
}

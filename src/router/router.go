package router

import (
	"gin-framework/basic/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//设置路由接口
func SetRouter(r *gin.Engine) *gin.Engine {
	r.Static("/static", "./public")           //设置静态资源
	r.StaticFS("/upload", http.Dir("upload")) //设置上传的资源
	r.Use(middleware.SetRouterGuard())        //设置路由监测
	api := r.Group("/api")
	UserRouter(api) //设置对应的User相关的路由设置
	return r
}

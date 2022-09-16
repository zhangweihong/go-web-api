package router

import (
	"gin-framework/basic/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) *gin.Engine {

	r.Static("/static", "./public")
	r.Use(middleware.SetRouterGuard())
	UserRouter(r)
	return r
}

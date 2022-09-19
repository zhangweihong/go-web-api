package router

import (
	user_controller "gin-framework/basic/src/controller/user_controller"

	"github.com/gin-gonic/gin"
)

//用户的路由
func UserRouter(r *gin.RouterGroup) {
	user := r.Group("user")
	user.GET("/", user_controller.FindeAllUser)
	user.POST("/avatar", user_controller.UploadAvatar)
}

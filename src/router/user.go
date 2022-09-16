package router

import (
	user_controller "gin-framework/basic/src/controller/user_controller"

	"github.com/gin-gonic/gin"
)

//用户的路由
func UserRouter(r *gin.Engine) {
	r.GET("/admin", user_controller.FindeAllUser)
	r.POST("/admin/avatar", user_controller.UploadAvatar)
}

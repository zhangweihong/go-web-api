package middleware

import "github.com/gin-gonic/gin"

func SetRouterGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ctx.Request.URL 进行url
		url := ctx.Request.URL.Path
		switch url {
		case "api/users/login":
			//进行鉴权等等
		}
	}
}

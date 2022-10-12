package controller

import (
	"gin-framework/basic/src/common"
	"gin-framework/basic/src/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

//返回前端数据
func toNet(ctx *gin.Context, code int, data any, message string) {
	backMap := make(map[string]any)
	backMap["status"] = code
	if data != nil {
		backMap["data"] = data
	}
	if message != "" {
		backMap["message"] = message
	}
	ctx.JSON(http.StatusOK, backMap)
}

//返回
func ToSuccess(ctx *gin.Context, data any) {
	toNet(ctx, 200, data, "")
}

//失败返回
func ToFail(ctx *gin.Context, code int, message string) {
	toNet(ctx, code, nil, message)
}

//验证参数函数
func ValidateParams(ctx *gin.Context, params any) bool {
	if err := ctx.ShouldBind(params); err != nil {
		ToFail(ctx, status.CommonError, err.Error()+" 参数验证错误")
		common.Logger.Warn(err.Error())
		return false
	}
	return true
}

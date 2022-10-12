package service

import (
	"gin-framework/basic/src/common"
	"gin-framework/basic/src/controller"
	db "gin-framework/basic/src/database"
	"gin-framework/basic/src/status"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Result struct {
	State int
	Data  any
	Err   string
}

//生成一个新的token
func NowToken(id string) string {
	var nowTokenKey = common.TokenKey(id)
	nowTokenChan := make(chan *redis.StringCmd)
	go func() {
		nowTokenChan <- db.RedisClient.Get(nowTokenKey)
	}()
	value, _ := (<-nowTokenChan).Result()
	defer close(nowTokenChan)
	return value
}

//设置一个token
func SetToken(token string, id string) string {
	var nowTokenKey = common.TokenKey(id)
	nowTokenChan := make(chan *redis.StatusCmd)

	go func() {
		nowTokenChan <- db.RedisClient.Set(nowTokenKey, token, time.Second*3600*24)
	}()
	status, _ := (<-nowTokenChan).Result()
	defer close(nowTokenChan)
	return status
}

//对比token
func CheckTokenIsValid(id string, checkToken string) bool {
	if checkToken == "" || id == "" {
		return false
	}
	var nowToken = NowToken(id)
	return nowToken == checkToken

}

//保存图片到本地
func SaveFile(ctx *gin.Context, file *multipart.FileHeader, savePath string) *Result {
	res_chan := make(chan *Result)
	go func() {
		err := ctx.SaveUploadedFile(file, savePath)
		if err != nil {
			common.Logger.Error(err.Error())
			res_chan <- &Result{
				State: -10003,
				Err:   err.Error(),
			}
			controller.ToFail(ctx, status.FileError, err.Error())
			return
		}
		res_chan <- &Result{
			State: 1,
		}
	}()
	return <-res_chan
}

package admin_controller

import (
	"fmt"
	"gin-framework/basic/src/helper"
	"gin-framework/basic/src/middleware"
	"gin-framework/basic/src/model"
	user_service "gin-framework/basic/src/service/user_service"
	"net/http"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
)

func FindeAllUser(ctx *gin.Context) {
	type Params struct {
		Test  string `form:"test" json:"test" binding:"required"`
		Test2 string `form:"test2" json:"test2" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		middleware.Logger.Error(err.Error())
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	chan_users := make(chan []model.User)
	go func() {
		chan_users <- user_service.SelectUser(ctx)
	}()
	ctx.JSON(http.StatusOK, gin.H{
		"user": <-chan_users,
	})
}

func UploadAvatar(ctx *gin.Context) {
	var wg sync.WaitGroup

	form, _ := ctx.MultipartForm()

	fmt.Println("form", form.File)
	files := form.File["avatar"]
	for _, file := range files {
		_file := file
		wg.Add(1)
		go func() {
			dst := helper.GetUploadsFilePath(_file.Filename)
			err := ctx.SaveUploadedFile(_file, dst)
			if err != nil {
				middleware.Logger.Error(err)
				wg.Done()
				runtime.Goexit()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	ctx.JSON(http.StatusOK, gin.H{
		"state": 1,
	})
}

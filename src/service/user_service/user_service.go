package admin_service

import (
	db "gin-framework/basic/src/database"
	"gin-framework/basic/src/model"

	"github.com/gin-gonic/gin"
)

func SelectUser(ctx *gin.Context) []model.User {

	chan_admins := make(chan []model.User)
	go func() {
		var users []model.User
		db.MysqlDb.Table("users").Find(&users)
		chan_admins <- users
	}()
	// s, err := db.RedisClient.Set("test_go_redis", 1, -1).Result()
	// fmt.Println(s, err)
	// v, err := db.RedisClient.Get("test_go_redis").Result()
	// v1, _ := strconv.Atoi(v)
	// fmt.Println(v, err, reflect.TypeOf(v1))
	defer close(chan_admins)
	return <-chan_admins
}

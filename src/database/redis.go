package db

import (
	"fmt"
	"gin-framework/basic/src/config"
	"gin-framework/basic/src/middleware"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

//初始化redis
func InitRedis() *redis.Client {
	redisConfig := config.Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password:    redisConfig.Password,     // no password set
		DB:          int(redisConfig.DBIndex), // use default DB
		PoolSize:    int(redisConfig.PoolNum),
		MaxRetries:  int(redisConfig.RetryNum),
		IdleTimeout: redisConfig.TimeOut,
	})

	_, err := RedisClient.Ping().Result()
	if err == nil {
		return RedisClient
	}
	middleware.Logger.Error(err)
	return nil
}

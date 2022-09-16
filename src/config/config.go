package config

import (
	"time"
)

var ServerPort = ":8080"

type MysqlConfig struct {
	Port            uint
	MaxConnsNum     uint
	MaxOpenConnsNUm uint
	User            string
	DbName          string
	Password        string
	Host            string
}

var Mysql *MysqlConfig = &MysqlConfig{
	Port:            3306,
	MaxConnsNum:     200,
	MaxOpenConnsNUm: 10,
	User:            "testUser",
	DbName:          "testDb",
	Password:        "testpwd123",
	Host:            "127.0.0.1",
}

type RedisConfig struct {
	Port     uint
	Host     string
	DBIndex  uint
	PoolNum  uint
	RetryNum uint
	TimeOut  time.Duration
	Password string
}

var Redis *RedisConfig = &RedisConfig{
	Port:     6379,
	Host:     "127.0.0.1",
	DBIndex:  0,
	PoolNum:  10,
	RetryNum: 5,
	TimeOut:  time.Duration(20 * time.Second),
	Password: "",
}

package db

import (
	"encoding/json"
	"fmt"
	"gin-framework/basic/src/config"
	"gin-framework/basic/src/middleware"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDb *gorm.DB

//初始化mysql
func InitMysql() *gorm.DB {
	var mysqlDbErr error
	mysqlConfig := config.Mysql
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DbName)
	MysqlDb, mysqlDbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if mysqlDbErr != nil {
		middleware.Logger.Error(mysqlDbErr)
	}
	db, err := MysqlDb.DB()
	if err != nil {
		middleware.Logger.Error(err)
	}
	db.SetMaxIdleConns(int(mysqlConfig.MaxConnsNum))     //最大连接数
	db.SetMaxOpenConns(int(mysqlConfig.MaxOpenConnsNUm)) //最大保持链接数（空闲数量）
	status, _ := json.Marshal(db.Stats())
	fmt.Printf("sql status %v \n", string(status))

	return MysqlDb
}

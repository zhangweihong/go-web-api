package db

func ConnectDb() {
	go InitMysql()
	go InitRedis()
}

package model

type User struct {
	Account  string `gorm:"primarykey;column:account"`
	Name     string `gorm:"primarykey;column:name"`
	Password string `gorm:"primarykey;column:password"`
}

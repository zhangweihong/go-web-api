package model

//对应mysql的users表
type User struct {
	Account  string `gorm:"primarykey;column:account" json:"account"`
	Name     string `gorm:"primarykey;column:name" json:"name"`
	Password string `gorm:"primarykey;column:password" json:"password"`
}

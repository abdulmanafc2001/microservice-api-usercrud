package models

type User struct {
	Id       int `gorm:"primaryKey"`
	UserName string
	Password string
	FullName string
	Phone    string
	Gender   string
}

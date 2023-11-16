package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Gender   string `json:"Gender"`
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
	Email string `json:"email" form:"email" query:"email"`
}

type Response struct {
	Message string
	Data interface{}
}
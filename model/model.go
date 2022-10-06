package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" query:"username" validate:"required,min=6"`
	Password string `json:"password" form:"password" query:"password" validate:"required,min=6,max=10"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
}

type Response struct {
	Message string
	Data interface{}
}

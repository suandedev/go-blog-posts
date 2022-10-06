package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	// migrate
	db.AutoMigrate(&User{})

	return db
}

// hash
func (user *User) Hash(password string) error{
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	user.Password = string(hash)
	return nil
}
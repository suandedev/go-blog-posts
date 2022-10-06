package main

import (
	"go-a/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func Router() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", GetUsers)
	e.POST("/users", PostUser)
	e.GET("/users/:username", GetUSer)
	e.DELETE("/users/:id", DeleteUser)
	e.PUT("/users/:id", UpdateUser)
	e.Any("/usr", GetData)
	return e
}

func GetUsers(c echo.Context) error {
	db := model.ConnectDb()

	var users []model.User
	db.Find(users)
	data := model.Response{Message: "oke", Data: users}
	return c.JSON(http.StatusOK, data)
}

func PostUser(c echo.Context) error {
	db := model.ConnectDb()

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}

	// hash
	if err := user.Hash(user.Password); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "error", Data: err.Error()})
	}
	db.Create(&user)
	data := model.Response{Message: "success", Data: user}
	return c.JSON(http.StatusOK, data)
}

func GetUSer( c echo.Context) error {
	username := c.Param("username")

	var user model.User
	// conn
	db := model.ConnectDb()
	db.Model(&user).Where("username = ?", username).First(&user)
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: user})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	// conn
	db := model.ConnectDb()
	db.Delete(&user, id)
	return c.JSON(http.StatusOK, model.Response{Message: "succes delete", Data: id})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	// con
	db := model.ConnectDb()

	// data
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}

	// hash
	if err := user.Hash(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: "error", Data: err.Error()})
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// update
	db.Model(&user).Where("id = ?", id).Updates(&user)
	
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: user})
}

func GetData(c echo.Context) (err error) {
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	return c.JSON(http.StatusOK, u)
}
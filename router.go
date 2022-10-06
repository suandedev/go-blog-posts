package main

import (
	"fmt"
	"go-a/model"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)


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
	// err
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", 
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", 
						err.Field())
				case "min":
					report.Message = fmt.Sprintf("%s value must be minim than %s",
						err.Field(), err.Param())
				case "max":
					report.Message = fmt.Sprintf("%s value must be max than %s",
						err.Field(), err.Param())
				}
	
				break
			}
		}
	
		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}
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

	var validate *validator.Validate
	
	validate = validator.New()
	if err := validate.Struct(user); err != nil {
		// return c.JSON(http.StatusBadGateway, err.Error())
		if caseObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range caseObject {
				switch err.Tag() {
				case "required":
					return c.JSON(http.StatusBadGateway, "is required " + err.Field())
				case "email":
					return c.JSON(http.StatusBadGateway, "is not valid email " + err.Field())
				case "min":
					return c.JSON(http.StatusBadGateway, "value " + err.Field() + " must more then " + err.Param())
				case "max":
					return c.JSON(http.StatusBadGateway, "value " + err.Field() + " must lower then " + err.Param())
				}
			}
		}
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
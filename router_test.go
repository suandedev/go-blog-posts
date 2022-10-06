package main

import (
	"go-a/model"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	r := gofight.New()
	r.GET("/").
		Run(Router(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			t.Log(r.Body.String())
			assert.Equal(t, "Hello, World!", r.Body.String())
			assert.Equal(t, 200, r.Code)
		})
}

func TestGetUsers(t *testing.T) {
	r := gofight.New()
	r.GET("/users").
		Run(Router(), func(h1 gofight.HTTPResponse, h2 gofight.HTTPRequest) {
			assert.Equal(t, 200, h1.Code)
		})
}

func TestPostUser(t *testing.T) {
	r := gofight.New()
	r.POST("/users").
		SetJSON(gofight.D{
			"username": "ssuande",
			"password": "secrets",
			"email": "made@gmail.com",
		}).
		Run(Router(), func(h1 gofight.HTTPResponse, h2 gofight.HTTPRequest) {
			assert.Equal(t, 200, h1.Code)
		})
}

// test hash
func TestHash(t *testing.T) {
	var user model.User
	user = model.User{
		Username: "made",
		Password: "secure",
		Email: "made@gmail.com",
	}
	hash := user.Hash(user.Password)
	assert.Equal(t, nil, hash, "error" )
}

// test get user
func TestGetUSer(t *testing.T) {
	r := gofight.New()
	r.GET("/users/mades").
		Run(Router(), func(h1 gofight.HTTPResponse, h2 gofight.HTTPRequest) {
			assert.Equal(t, 400, h1.Code, "not success" + h1.Body.String())
		})
}

// test delete user
func TestDeleteUSer(t *testing.T) {
	r := gofight.New()
	r.DELETE("/users/3").
		Run(Router(), func(h1 gofight.HTTPResponse, h2 gofight.HTTPRequest) {
			assert.Equal(t, 200, h1.Code, "not match")
		})
}

// test update user
func TestUpdateUser(t *testing.T) {
	r := gofight.New()
	r.PUT("/users/1").
		SetForm(gofight.H{
			"username": "made2100",
			"password": "2secret2",
			"email": "made@gmail.com",
		}).
		Run(Router(), func(h1 gofight.HTTPResponse, h2 gofight.HTTPRequest) {
			assert.Equal(t, 200, h1.Code, "not match" )
		})
}

// get data
func TestGetData(t *testing.T) {
	r := gofight.New()
	r.POST("/usr").
		SetJSON(gofight.D{
			"username": "made",
			"password" : "secret",
			"email": "made@gmail.com",
		}).
		Run(Router(), func(h1 gofight.HTTPResponse, h2 gofight.HTTPRequest) {
			assert.Equal(t, 201, h1.Code, "not match" + h1.Body.String())
		})
}
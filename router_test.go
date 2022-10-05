package main

import (
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
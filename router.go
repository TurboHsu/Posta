package main

import (
	h "github.com/TurboHsu/Posta/handler"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", h.Ping)
	r.GET("/q", h.QueryTest)
	return r
}

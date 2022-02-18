package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	r := gin.Default()
	r.POST("/register", register)
	r.POST("/login", login)

	r.Run()
}
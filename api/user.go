package api

import (
	"Douban/model"
	"Douban/service"
	"Douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	r := gin.Default()
	username := c.PostForm("username")
	password := c.PostForm("password")

	flag, err := service.IsPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password err: ", err)
		tool.RespInternalError(c)
		return
	}

	if !flag {
		tool.RespErrorWithDate(c, "密码错误")
		return
	}

	r.GET("/login", func(c *gin.Context) {
		userCookie := &http.Cookie{
			Name:     "username",
			Value:    username,
			Path:     "/login",
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, userCookie)
	})

	tool.RespSuccessful(c)
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	flag, err := service.IsRepeatUsername(username)
	if err != nil {
		fmt.Println("judge repeat username err: ", err)
		tool.RespInternalError(c)
		return
	}

	if flag {
		tool.RespErrorWithDate(c, "用户名重复")
		return
	}

	err = service.Register(user)
	if err != nil {
		fmt.Println("register err: ", err)
		tool.RespInternalError(c)
		return
	}

	tool.RespSuccessful(c)
}

func userQuit(c *gin.Context) {
	cookieKey := "username"
	_, err := c.Request.Cookie(cookieKey)
	if err != nil {
		tool.RespInternalError(c)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "已登出",
		})
	}
}

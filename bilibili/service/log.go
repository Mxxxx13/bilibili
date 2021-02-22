package service

import (
	"bilibili/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取用户注册的信息
func Register(ctx *gin.Context) bool {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	res := models.Register(username, password)
	return res
}

//获取用户登录的信息
func Login (ctx *gin.Context) bool {
	cookie1, err1 := ctx.Request.Cookie("username")
	cookie2, err2 := ctx.Request.Cookie("password")
	if err1 == nil && err2 == nil {
		res := models.Login(cookie1.Value, cookie2.Value)
		return res
	} else {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		bl := ctx.PostForm("bool")
		res := models.Login(username, password)
		if res == true {
			Username := &http.Cookie{
				Name : "username",
				Value: username,
				Path: "/",
				MaxAge: 60,
				HttpOnly: true,
			}
			http.SetCookie(ctx.Writer, Username)
		}
		//账号密码匹配且记住密码时设置cookie
		if res == true && bl == "true" {
			Password := &http.Cookie{
				Name : "password",
				Value: password,
				Path: "/",
				MaxAge: 60,
				HttpOnly: true,
			}
			http.SetCookie(ctx.Writer, Password)
		}
		return res
	}
}

//func Logout(ctx *gin.Context) bool {
//	return true
//}
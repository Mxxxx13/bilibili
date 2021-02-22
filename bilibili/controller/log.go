package controller

import (
	"bilibili/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册
func Register(ctx *gin.Context){
	res := service.Register(ctx)
	if res {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 100,
			"message" : "注册成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"message" : "注册失败",
		})
	}
}

//登录
func Login(ctx *gin.Context){
	res := service.Login(ctx)
	if res == true {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 100,
			"message" : "登录成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"message" : "登录失败",
		})
	}
}

////登出
//func Logout(ctx *gin.Context){
//	res := service.Logout(ctx)
//	if res {
//		ctx.JSON(http.StatusOK, gin.H{
//			"code" : 100,
//			"message" : "注销成功",
//		})
//	} else {
//		ctx.JSON(http.StatusOK, gin.H{
//			"code" : 200,
//			"message" : "注销失败",
//		})
//	}
//}


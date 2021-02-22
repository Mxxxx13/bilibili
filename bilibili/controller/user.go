package controller

import (
	"bilibili/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//个人空间
func Space(ctx *gin.Context) {
	u,res := service.Space(ctx)
	if res == true {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 100,
			"message" : "成功",
		})
		ctx.JSON(http.StatusOK,u)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"message" : "失败",
		})
	}
}
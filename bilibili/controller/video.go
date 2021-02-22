package controller

import (
	"bilibili/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//视频投稿
func PostVideo(ctx *gin.Context){
	res := service.PostVideo(ctx)
	if res {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 100,
			"message" : "投稿成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"message" : "投稿失败",
		})
	}
}

//视频详情
func ShowVideo(ctx *gin.Context) {
	u,res := service.ShowVideo(ctx)
	if res  {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 100,
			"message" : "成功",
		})
		ctx.JSON(http.StatusOK,u)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"message" : "视频不存在",
		})
	}
}

//搜索视频
func Search(ctx *gin.Context) {
	res := service.Search(ctx)
	if res  {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 100,
			"message" : "成功",
		})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"message" : "失败",
		})
	}
}

//删除视频
func DeleteVideo(ctx *gin.Context) {
	res := service.DeleteVideo(ctx)
	if res  {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 100,
			"message" : "视频删除成功",
		})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"message" : "视频删除失败",
		})
	}
}
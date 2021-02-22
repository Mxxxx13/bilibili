package service

import (
	"bilibili/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostVideo(ctx *gin.Context) bool {
	var video models.Video
	if err := ctx.ShouldBind(&video); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Title": video.Title,
			"Image": video.Image,
			"Info": video.Info,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	cookie , err1 :=  ctx.Request.Cookie("username")
	if err1 == nil {
		video.Author = cookie.Value
	}
	res := models.PostVideo(video.Title,video.Image,video.Author,video.Info)
	return res
}

func ShowVideo(ctx *gin.Context) (video models.Video,b bool) {
	str := ctx.Param("id")
	id, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		fmt.Printf("failed:%v\n", err)
	}
	video,bool := models.ShowVideo(int(id))
	return video,bool
}

func Search(ctx *gin.Context)  bool {
	str := ctx.Query("keyword")
	res := models.Search(str,ctx)
	return res
}

func DeleteVideo(ctx *gin.Context)  bool {
	str := ctx.Param("id")
	id, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		fmt.Printf("failed:%v\n", err)
	}
	cookie,err := ctx.Request.Cookie("username")
	if err != nil {
		fmt.Printf("get cookie failed %v",err)
	} else {
		username:= cookie.Value
		res := models.DeleteVideo(int(id),username)
		return res
	}

	return false
}

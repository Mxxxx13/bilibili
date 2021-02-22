package cmd

import (
	"bilibili/controller"
	"bilibili/middleware"
	"github.com/gin-gonic/gin"
)

func Entrance(){
	router := gin.Default()
	//登录注册界面
	log := router.Group("passport/bilibili")
	{
		log.POST("/register",controller.Register)
		log.GET("/login",controller.Login)
	}

	bili := router.Group("bilibili")
	{

		bili.GET("/video/:id",controller.ShowVideo)
		bili.GET("/all",controller.Search)
	}
	//需要登录保护
	auth := router.Group("/")
	auth.Use(middleware.AuthRequired())
	{
		auth.GET("/space/bilibili/:id", controller.Space)
		//auth.DELETE("/passport/bilibili/logout",controller.Logout)
		v := auth.Group("/bilibili/video")
		{
			v.POST("/post",controller.PostVideo)
			v.DELETE("/:id",controller.DeleteVideo)
		}
	}
	router.Run(":8080")
}
package service

import (
	"bilibili/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Space(ctx *gin.Context) (user models.User,b bool) {
	str := ctx.Param("id")
	id, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		fmt.Printf("failed:%v\n", err)
	}
	u,bool := models.Space(int(id))
	return u,bool
}
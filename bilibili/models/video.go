package models

import (
	"bilibili/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Video struct {
	Id int 			`json:"id"`
	Title string 	`json:"title" binding:"required,min=2,max=16"`
	Image string 	`json:"image"`
	Author string 	`json:"author"`
	Info string 	`json:"info" binding:"required,min=0,max=200"`
	Views int 		`json:"views" `
	Likes int 		`json:"likes"`
	Coins int 		`json:"coins"`
	Collections int `json:"collections"`
}

func PostVideo(Title string,Image string , Author string, Info string) bool{
	stmt, err := dao.DB.Prepare("insert into video(Title,Image,Author,Info) value (?,?,?,?)")
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(Title, Image, Author, Info)
	if err != nil {
		fmt.Printf("insert failed:%v", err)
		return false
	}
	return true
}

func ShowVideo(id int) (video Video,b bool){
	var v Video
	stmt, err := dao.DB.Query("select * from video where id = ?",id)
	if err != nil {
		fmt.Printf("query failed:%v", err)
		return  v,false
	}
	defer stmt.Close()
	for stmt.Next() {
		err = stmt.Scan(&v.Id,&v.Title,&v.Image,&v.Author,&v.Info,&v.Views,&v.Likes,&v.Coins,&v.Collections)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return  v,false
		} else {
			return v,true
		}
	}
	return  v,false
}

func Search(keyword string,ctx *gin.Context)  bool {
	var v Video
	var count int
	stmt, err := dao.DB.Query("select * from video where title like concat('%',?,'%') or author like concat('%',?,'%')",keyword,keyword)
	if err != nil {
		fmt.Printf("query failed:%v", err)
		return  false
	}
	defer stmt.Close()
	for stmt.Next() {
		err = stmt.Scan(&v.Id,&v.Title,&v.Image,&v.Author,&v.Info,&v.Views,&v.Likes,&v.Coins,&v.Collections)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return  false
		} else {
			ctx.JSON(http.StatusOK,v)
			count++
		}
	}
	if count != 0 {
		return true
	}
	return  false
}

func DeleteVideo(id int,username string) bool{
	stmt, err := dao.DB.Prepare("delete from video where id = ? and author = ?")
	if err != nil {
		fmt.Printf("Prepare failed:%v", err)
		return  false
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, username)
	if err != nil {
		fmt.Printf("delete failed:%v", err)
		return false
	}
	return true
}
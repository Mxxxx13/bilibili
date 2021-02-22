package models

import (
	"bilibili/dao"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int 			`json:"id"`
	Username string `json:"username" binding:"required,min=2,max=16"`
	Password string `json:"password" binding:"required,min=6,max=16"`
	Level int 		`json:"level"`
	Subscribe int 	`json:"subscribe"`
	Fans int 		`json:"fans"`
	Likes int		`json:"likes"`
	Plays int		`json:"plays"`
	Views int 		`json:"views"`
}

func Register(username string, password string) bool{
	stmt, err := dao.DB.Prepare("insert into user(username,password) value (?,?)")
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil {
		fmt.Printf("insert failed:%v", err)
		return false
	}
	return true
}

func Login(username string, password string) bool{
	var u User
	stmt, err := dao.DB.Query("select password from user where username=?", username)
	if err != nil {
		fmt.Printf("query failed:%v", err)
		return false
	}
	defer stmt.Close()
	for stmt.Next() {
		err = stmt.Scan(&u.Password)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return false
		}
	}
	if u.Password == password && password != ""{
		return true
	}
	return false
}


//获取登录用户的id并记录登录状态
func GetUserId(username string) int {
	var u User
	stmt, err := dao.DB.Query("select id from user where username=?", username)
	if err != nil {
		fmt.Printf("query failed:%v", err)
		return 0
	}
	defer stmt.Close()
	for stmt.Next() {
		err = stmt.Scan(&u.Id)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return 0
		}
		return u.Id
	}
	return 0
}
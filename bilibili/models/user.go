package models

import (
	"bilibili/dao"
	"fmt"
)

func Space(id int) (user User,b bool){
	var u User
	stmt, err := dao.DB.Query("select * from user where id = ?",id)
	if err != nil {
		fmt.Printf("query failed:%v", err)
		return  u,false
	}
	defer stmt.Close()
	for stmt.Next() {
		err = stmt.Scan(&u.Id,&u.Username,&u.Password,&u.Level,&u.Subscribe,&u.Fans,&u.Likes,&u.Plays,&u.Views)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return  u,false
		} else {
			return u,true
		}
	}
	return  u,false
}

package main

import (
	"bilibili/cmd"
	"bilibili/dao"
)

func main() {
	dao.MysqlInit()
	cmd.Entrance()
}

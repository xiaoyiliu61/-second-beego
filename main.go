package main

import (
	"beefile/db_mysql"
	_ "beefile/routers"
	"github.com/astaxie/beego"
)

func main() {
    //1.连接数据库
	db_mysql.Connect()

	beego.Run()
}


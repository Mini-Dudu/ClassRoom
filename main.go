package main

import (
	"Beego_Web/database"
	_ "Beego_Web/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	//打开数据库
	str,err := database.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer database.CloseDb()

	fmt.Println(str)
	beego.Run()
}

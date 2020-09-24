package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//获得请求参数
	name := c.Ctx.Input.Query("name")
	password := c.Ctx.Input.Query("password")

	if name == "dudu" && password == "123456" {
		c.Ctx.ResponseWriter.Write([]byte("欢迎来到嘟嘟的主页"))
	}else {
		c.Ctx.ResponseWriter.Write([]byte("用户名或者密码错误，请重试！"))
	}

	fmt.Println("用户名：",name,"密码:",password)
}

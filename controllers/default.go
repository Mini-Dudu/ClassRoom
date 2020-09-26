package controllers

import (
	"Beego_Web/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

//get方法的重写
func (c *MainController) Get() {
	//获得请求参数
	name := c.Ctx.Input.Query("name")
	password := c.Ctx.Input.Query("password")

	if name != "dudu" || password != "123456" {
		c.Ctx.ResponseWriter.Write([]byte("用户名或者密码错误，请重试！"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("欢迎来到嘟嘟的主页"))
	fmt.Println("用户名：",name,"密码:",password)
}


////处理普通文本类型的post
//func (c *MainController) Post() {
//	//获得请求参数
//	name := c.Ctx.Request.FormValue(("name"))
//	password := c.Ctx.Request.FormValue(("password"))
//
//	//fmt.Printf("%T   %T\n",name,password)
//	//fmt.Printf("%t   %t",name == "name",password == "123456")
//
//	if name != "dudu" || password != "123456" {
//		c.Ctx.WriteString("用户名或者密码错误，请重试！")
//		return
//	}
//
//	c.Ctx.WriteString("欢迎来到嘟嘟的主页")
//
//	fmt.Println("用户名：",name,"密码:",password)
//}


//处理JSON类型的post
func (c *MainController) Post() {

	//解析JSON格式数据
	var dudu models.Dudu_info

	dataBytes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收出错，请重试！")
		return
	}

	fmt.Println(string(dataBytes))

	err = json.Unmarshal(dataBytes,&dudu)

	if err != nil {
		fmt.Println("错误",err.Error())
		c.Ctx.WriteString("解析数据出错，请重试！")
		return
	}

	c.Ctx.WriteString("数据解析成功！")
	//fmt.Printf("name:%s age:%s hobby: %s",user.Name,user.Password)

}

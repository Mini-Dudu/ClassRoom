package controllers

import (
	"Beego_Web/database"
	"Beego_Web/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strconv"
)

type UserLogin struct {
	beego.Controller
}


//处理用户登入时JSON格式数据的post请求
func (login *UserLogin) Post(){

	bodyDatas,err := ioutil.ReadAll(login.Ctx.Request.Body)
	if err != nil {
		login.Ctx.WriteString("数据接收出错，请重试！")
		return
	}

	//2.将数据保存到结构体中
	var user models.UserLogin
	err = json.Unmarshal(bodyDatas,&user)
	if err != nil {
		login.Ctx.WriteString("数据保存至user中出错，请重试！")
		return
	}

	result,_ := database.User_login(strconv.Itoa(user.UserId), user.PassWord)
	login.Ctx.WriteString(result)
}
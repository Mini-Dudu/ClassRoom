package controllers

import (
	"Beego_Web/database"
	"Beego_Web/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
)

//处理register请求
type Register struct {
	beego.Controller
}


//处理用户注册时JSON格式数据的post请求
func (r *Register) Post(){

	//1.接收前端传递的数据
	//方法一:
	//bodyDatas,err := r.Ctx.Request.GetBody()
	//方法二
	bodyDatas,err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接收出错，请重试！")
		return
	}

	//2.将数据保存到结构体中
	var user models.User
	err = json.Unmarshal(bodyDatas,&user)
	if err != nil {
		r.Ctx.WriteString("数据保存至user中出错，请重试！")
		return
	}

	//3.将解析后得到的user数据保存到数据库中
	_,err = database.Inssert_user(user)
	if err != nil {
		 r.Ctx.WriteString("将解析后得到的user数据保存到数据库中时出错，请重试！")
		return
	}

	//4.反馈注册成功与否
	//方法一:
	r.Ctx.WriteString("注册成功")

	////方法二:
	//result := models.ResponseResult{
	//	Code:    0,
	//	Message: "注册成功",
	//	Data:    nil,
	//}
	//
	//r.Data["json"] = &result
	//r.ServeJSONP()
}

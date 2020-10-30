package routers

import (
	"Beego_Web/controllers"
	"github.com/astaxie/beego"
)

func init() {

	//默认处理
    beego.Router("/", &controllers.MainController{})

	//处理注册用户
	beego.Router("/reginster",&controllers.Register{})

	//处理用户登入
	beego.Router("/login",&controllers.UserLogin{})
	
	
}

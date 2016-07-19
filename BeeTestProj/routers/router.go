package routers

import (
	"GoLang-PlayGround/BeeTestProj/controllers"
	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})
}

package routers

import (
	"GoLang-PlayGround/BeeTestProj/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

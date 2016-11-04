package routers

import (
	"GoLang-PlayGround/BeegoPlayGround/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

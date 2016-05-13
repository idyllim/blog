package routers

import (
	"github.com/astaxie/beego"
	"idyllim.com/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

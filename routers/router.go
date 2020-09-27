package routers

import (
	"beefile/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register", &controllers.RegisterController{})

	beego.Router("/login", &controllers.MainController{})
}

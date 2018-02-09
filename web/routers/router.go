package routers

import (
	"github.com/astaxie/beego"
	"github.com/restic/restic/web/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
}

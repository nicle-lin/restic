package main

import (
	"github.com/astaxie/beego"
	_ "github.com/restic/restic/web/routers"
	"github.com/restic/restic/web/controllers"
	"github.com/astaxie/beego/orm"
)

func main() {
	err := InitMysqlDB()
	if err != nil {
		panic(err)
	}
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}


	beego.SetStaticPath("/img","static/img")
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/js","static/js")
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

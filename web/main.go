package main

import (
	"github.com/astaxie/beego"
	_ "github.com/restic/restic/web/routers"
)

func main() {
	beego.SetStaticPath("/img","static/img")
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/js","static/js")
	beego.Run()
}

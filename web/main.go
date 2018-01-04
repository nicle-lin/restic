package main

import (
	"github.com/astaxie/beego"
	_ "github.com/restic/restic/web/routers"
)

func main() {
	beego.Run()
}

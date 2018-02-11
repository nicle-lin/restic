package controllers

import (
	"github.com/astaxie/beego"
)


type BaseController struct {
	beego.Controller
}


type GrantController struct {
	BaseController
}

func (c *BaseController)Flash(cc *beego.Controller){
	flash := beego.ReadFromRequest(cc)
	_, ok1 := flash.Data["error"]
	_, ok2 := flash.Data["warning"]
	_, ok3 := flash.Data["notice"]
	if ok1 || ok2 || ok3 {
		cc.Data["Error"] = true
	}
}

func(c *BaseController)Prepare(){

}

func(c *BaseController)Finish(){

}


type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "error/404.html"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "error/501.html"
}


func (c *ErrorController) ErrorDatabase() {
	c.Data["content"] = "database is now down"
	c.TplName = "error/database.html"
}
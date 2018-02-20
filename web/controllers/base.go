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
	_, ok := flash.Data["error"]
	if ok{
		cc.Data["FlashError"] = true
		return
	}
	_, ok = flash.Data["warning"]
	if ok{
		cc.Data["FlashWarning"] = true
		return
	}
	_, ok = flash.Data["notice"]
	if ok{
		cc.Data["FlashNotice"] = true
		return
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
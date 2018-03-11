package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	beego.Controller
	flash *beego.FlashData
	valid validation.Validation
}

type GrantController struct {
	BaseController
}

func (c *BaseController) Flash(cc *beego.Controller) {
	flash := beego.ReadFromRequest(cc)
	_, ok := flash.Data["error"]
	if ok {
		cc.Data["FlashError"] = true
		return
	}
	_, ok = flash.Data["warning"]
	if ok {
		cc.Data["FlashWarning"] = true
		return
	}
	_, ok = flash.Data["notice"]
	if ok {
		cc.Data["FlashNotice"] = true
		return
	}
}
func (c *BaseController) ErrorDatabase(msg string) {
	if beego.BConfig.RunMode == "dev" {
		c.Data["error"] = msg
	} else {
		c.Data["error"] = "数据库异常,请联系我,邮箱:dghpgyss@163.com"
	}
	c.Abort("Database")
}

func (c *BaseController) FlashError(msg string, args ...interface{}) {
	if beego.BConfig.RunMode == "dev" {

	}
}

func (c *BaseController) Prepare() {
	c.flash = beego.NewFlash()
	c.valid = validation.Validation{}
	c.GetSession("email")
	c.GetSession("password")
}

func (c *BaseController) Finish() {

}

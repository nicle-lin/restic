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



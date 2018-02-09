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

func(c *BaseController)Prepare(){

}

func(c *BaseController)Finish(){

}

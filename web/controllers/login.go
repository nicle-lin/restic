package controllers

import (
	"github.com/astaxie/beego"
	"github.com/restic/restic/web/models"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.Flash(&c.Controller)
	c.TplName = "login/login.html"
	return
}

func (c *LoginController) Post() {
	type Params struct {
		Email      string `form:"email"`
		Password   string `form:"password"`
		RememberMe int64  `form:"remember_me"`
	}
	var p Params
	if err := c.ParseForm(&p); err != nil {
		panic(err)
	}

	flash := beego.NewFlash()
	models.UserByEmail(p.Email)
	if p.Email != "dghpgyss@163.com" || p.Password != "admin" {
		//c.Abort("Database")
		flash.Error("email or password is incorrect")
		flash.Store(&c.Controller)
		c.Redirect("/", 301) // redirect to home
		return
	}

	c.Ctx.SetCookie("email", p.Email, 24*3600, "/")
	c.Ctx.SetCookie("password", p.Password, 24*3600, "/")
	c.Redirect("/home", 301) // redirect to home
	return
}

package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/restic/restic/web/internal/common"
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

type Params struct {
	Email      string `form:"email"`
	Password   string `form:"password"`
	RememberMe int64  `form:"remember_me"`
}

func (c *LoginController) Post() {

	var p Params
	if err := c.ParseForm(&p); err != nil {
		beego.Debug(err)
	}

	u, err := models.UserByEmail(p.Email)
	if err != nil {
		c.ErrorDatabase(err.Error())
		return
	}
	//md5
	if u == nil || Md5sum(p.Password) != u.Password {
		c.flash.Notice("email or password is incorrect")
		c.flash.Store(&c.Controller)
		c.Redirect("/", 301) // redirect to login
		return
	}

	c.SetSession("email", u.Email)
	c.SetSession("password", Md5sum(u.Password))
	c.Redirect("/home", 301) // redirect to home
	return
}

package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/restic/restic/web/internal/common"
	"github.com/restic/restic/web/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	c.Flash(&c.Controller)
	c.TplName = "user/new_user.html"
	return
}

type NewUser struct {
	Email         string `form:"email" valid:"Required"`
	Password      string `form:"password" valid:"Required"`
	PasswordAgain string `form:"password_again" valid:"Required"`
}

func (c *UserController) Post() {
	var p Params
	if err := c.ParseForm(&p); err != nil {
		beego.Debug(err)
	}
	c.valid.Valid(&p)
	for _, err := range c.valid.Errors {
		c.flash.Error(err.Error())
		c.flash.Store(&c.Controller)
		c.Redirect("/user", 301)
		return
	}
	data := map[string]interface{}{
		"email":    p.Email,
		"password": Md5sum(p.Password),
	}
	models.NewUser(data)
	c.Redirect("/home", 301)
	return
}

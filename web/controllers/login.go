package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.Flash(&c.Controller)
	c.TplName = "login/login.html"
	return
}

type User struct {
	Email string `form:"email"`
	Password string `form:"password"`
	RememberMe int64 `form:"remember_me"`
}
func (c *LoginController)Post()  {
	flash := beego.NewFlash()
	var u User
	if err := c.ParseForm(&u); err != nil{
		panic(err)
	}
	if u.Email != "dghpgyss@163.com" || u.Password != "admin"{
		//c.Abort("Database")
		flash.Error("email or password is incorrect")
		flash.Store(&c.Controller)
		c.Redirect("/login",301) // redirect to home
		return
	}

	c.Ctx.SetCookie("email",u.Email,24*3600, "/")
	c.Ctx.SetCookie("password",u.Password,24*3600, "/")
	c.Redirect("/",301) // redirect to home
	return
}

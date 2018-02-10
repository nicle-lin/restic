package controllers

import "fmt"

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login/login.html"
}

type User struct {
	Email string
	Password string
}
func (c *LoginController)Post()  {
	var u User
	if err := c.ParseForm(&u); err != nil{
		panic(err)
	}
	c.Data["Password"] = u.Password
	c.Data["Email"] = u.Email
	c.TplName = "login/login.html"
	fmt.Println(u.Email)
	c.RenderString()
}

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
func (c *LoginController)Post()  {
	fmt.Println(c.Ctx.Input.RequestBody)

}

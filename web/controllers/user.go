package controllers

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	c.Flash(&c.Controller)
	c.TplName = "login/login.html"
	return
}

func (c *UserController) Post() {
	return
}

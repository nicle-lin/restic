package controllers


type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Password"] = "test"
	c.TplName = "home/home.html"
}

func (c *HomeController)Post()  {
}
package controllers


type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	c.TplName = "home/home.html"
}

func (c *HomeController)Post()  {
	c.TplName = "home/home.html"
}
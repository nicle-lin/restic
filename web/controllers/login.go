package controllers


type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login/login.html"
}

type User struct {
	Email string `form:"email"`
	Password string `form:"password"`
}
func (c *LoginController)Post()  {
	var u User
	if err := c.ParseForm(&u); err != nil{
		panic(err)
	}
	c.Data["Password"] = u.Password
	c.Data["Email"] = u.Email
	c.TplName = "home/home.html"
}

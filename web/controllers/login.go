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
	RememberMe int64 `form:"remember_me"`
}
func (c *LoginController)Post()  {
	var u User
	if err := c.ParseForm(&u); err != nil{
		panic(err)
	}
	if u.Email != "dghpgyss@163.com" || u.Password != "admin"{
		c.Abort("501")
		return
	}

	c.Ctx.SetCookie("email",u.Email,24*3600, "/")
	c.Ctx.SetCookie("password",u.Password,24*3600, "/")
	c.Redirect("/",301) // redirect to home
	return
}

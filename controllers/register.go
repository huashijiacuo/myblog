package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	beego.Informational("get请求注册页")
	c.Data["Website"] = "beego.me.shun"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "register.tpl"

	//c.Ctx.WriteString("hello")
}

func (r *RegisterController) Post() {
	beego.Informational("注册请求")
	u := models.User{}
	if err := r.ParseForm(&u); err != nil {
		//handle error
	}
	beego.Informational("username:", u.UserName, " Email:", u.Email, "sex:", u.Sex, "pwd:", u.PassWord)
	message, err := models.InsertUser(&u)
	beego.Informational(message)
	if err != nil {
		r.Redirect("/register", 301)
		r.Ctx.WriteString("failed to create the user!")
		beego.Error(err)

	}
	// r.Ctx.WriteString(message)

	r.Redirect("/", 301)
	//c.Ctx.WriteString("hello")
}

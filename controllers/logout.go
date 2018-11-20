package controllers

import "github.com/astaxie/beego"

type LogoutController struct {
	beego.Controller
}

// Logout 退出
func (this *LogoutController) Get() {
	//this.SetSecureCookie(Secret, Token, "", -1, Path, Domain, Secure, HttpOnly)
	this.Ctx.SetCookie("username", "-1", 100, "/")  // 设置cookie
	userName := this.Ctx.GetCookie("username")
	token := this.Ctx.GetCookie("token")
	if userName != "" && token != "-1" {
		beego.Informational("并没有登录，无需设置，直接返回首页！")
	} else {
		sessionToken := this.GetSession(userName)
		if sessionToken != "" {
			this.DelSession(userName)
		}
		this.Ctx.SetCookie("username", "")
		this.Ctx.SetCookie("token", "-1")
	}
	this.Redirect("/", 302)
}
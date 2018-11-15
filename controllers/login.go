package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}
// Logout 退出
func (this *LoginController) Logout() {
	//this.SetSecureCookie(Secret, Token, "", -1, Path, Domain, Secure, HttpOnly)
	this.Redirect("/auth/login", 302)
}
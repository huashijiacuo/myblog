package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"myblog/models"
	"strconv"
)
var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	fmt.Println("get请求")
	c.Data["Website"] = "beego.me.shun"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	//c.Ctx.WriteString("hello")
}

func (this *MainController) Post() {

	username := this.GetString("username")
	pwd := this.GetString("pwd")
	beego.Informational("登录请求！", "username:", username, " pwd:", pwd)

	name := this.Ctx.GetCookie("username")
	//password := this.Ctx.GetCookie("password")
	beego.Informational("cookie中的username:" + name)
	if name != "" {
		beego.Informational("cookie中已有值，已登录")
		this.Ctx.WriteString("已登录，无需再次登录!")
		return
	}

	u := models.GetUserByName(username)
	beego.Informational(u.PassWord)
	if u.PassWord == pwd {
		beego.Informational("帐号密码正确，允许登录，并设置cookie!")
		this.Ctx.SetCookie("username", u.UserName, 100, "/")  // 设置cookie
		this.Ctx.SetCookie("password", u.PassWord, 100, "/")  // 设置cookie
		this.Ctx.WriteString(u.UserName + "登录成功！用户信息写入cookie!")
		return
	}
	beego.Informational("输入密码有误，请重新登录！")
	//this.Redirect("/", 200)
	this.Ctx.WriteString(u.UserName + "登录失败，密码错误!")
}


type user struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}

func (u *user) string() string {
	if len(u.Name.(string)) != 0 {
		fmt.Println("u.Name is ok")
	}
	return "Id:" + strconv.Itoa(u.Id) + "\nName:" + u.Name.(string) + "\nAge:" + strconv.Itoa(u.Age) + "\nEmail:" + u.Email
}

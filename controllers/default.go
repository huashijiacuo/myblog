package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

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
	u := user{}
	fmt.Println("post请求！")
	if err := this.ParseForm(&u); err != nil {
		//handle error
	}
	fmt.Println(u)

	this.Ctx.WriteString(u.string())
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

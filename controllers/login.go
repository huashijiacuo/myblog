package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"myblog/models"
	"net/http"
)

type LoginController struct {
	beego.Controller
}

type UserFront struct {
	UserName string `form:"userName"`
	PassWord string `form:"userPwd"`
}

func (this *LoginController) Get() {
	beego.Informational("get请求login页面！！")
	body := this.Ctx.Input.RequestBody
	beego.Informational(body)
	this.Ctx.SetCookie("username", "")
	this.Ctx.SetCookie("token", "-1")
	this.TplName = "login.tpl"
}



func (this *LoginController) Login() {
	var err error
	var u UserFront//先声明err类型
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &u)
	if err == nil {
		//fmt.Println(r)大家到这里可以试一试自己有没有成功获取JSON数据，成功的话这里会输出我们需要的JSON字符串
		//content := u.(map[string]interface{})      //这里通过一个map里总体匹配r中的数据
		beego.Informational("前端传递过来的json数据：" + u.UserName + " " + u.PassWord)

		userJson, err := json.Marshal(u)
		if err != nil{
			panic(err)
		}

		//Set Content-Type header so that clients will know how to read response
		this.Ctx.ResponseWriter.Header().Set("Content-Type","application/json")
		this.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
		//Write json response back to response
		this.Ctx.ResponseWriter.Write(userJson)
		return
	}
	beego.Informational("自定义路由成功！")
	this.Ctx.WriteString("访问成功！")
}


func (this *LoginController) Post() {
	beego.Informational("login登录请求post!")
	var err error
	var u UserFront//先声明err类型
	body := this.Ctx.Input.RequestBody//这是获取到的json二进制数据
	beego.Informational("*****************body*************\n")
	beego.Informational(body)
	beego.Informational("*****************body*************\n")
	err = json.Unmarshal(body, &u)//解析二进制json，把结果放进ob中
	beego.Informational("************json数据*****************")
	if err == nil {
		//fmt.Println(r)大家到这里可以试一试自己有没有成功获取JSON数据，成功的话这里会输出我们需要的JSON字符串
		beego.Informational(u.UserName + " " + u.PassWord)

		this.Ctx.ResponseWriter.Header().Set("Content-Type","application/json")
		this.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
		//Write json response back to response
		this.Ctx.WriteString("转化成功！")
		return
	} else {
		beego.Informational("json转化失败！")
		err2 := this.ParseForm(&u)
		if err2 != nil {
			beego.Informational("json转化失败！")
		} else {
			beego.Informational("congratulation! json 转化成功！")
			beego.Informational(u.UserName + " " + u.PassWord)
			user := models.GetUserByName(u.UserName)
			if (user.PassWord == u.PassWord) {
				beego.Informational("Congratulation! 密码正确，登录成功,并设置cookie和session！")
				//this.Ctx.ResponseWriter.Header().Set("Content-Type","application/json")
				//this.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
				this.Ctx.SetCookie("username", u.UserName)
				this.Ctx.SetCookie("token", "1")
				this.SetSession(u.UserName, "1")
			} else {
				beego.Informational("登录失败！")
			}
		}
	}


	this.Ctx.WriteString("访问成功！")
}
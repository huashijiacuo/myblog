package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "myblog/models"
	_ "myblog/routers"
)




func main() {
	fmt.Println("This is my blog which developed by go!!")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Informational(beego.BConfig.CopyRequestBody)
	beego.SetStaticPath("/download", "static/download")
	beego.Run()

	//u := models.GetUserByName("shi")
	//fmt.Println(u)
	//
	//uid := models.GetUserById(3)
	//fmt.Println("%+v",uid)

}


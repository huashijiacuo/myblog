package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "myblog/models"
	_ "myblog/routers"
)

func main() {
	fmt.Println("This is my blog which developed by go!!")
	beego.SetStaticPath("/download", "static/download")
	beego.Run()

}


package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get() {
	this.TplName = "upload.tpl"
}

func (this *UploadController) Post() {
	f, h, err := this.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	this.SaveToFile("uploadname", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建

	this.Redirect("/",302)
}


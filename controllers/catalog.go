package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type CatalogController struct {
	beego.Controller
}

type FrontCatalog struct {
	CataName string
	CataUserName string
}


/**
 *get request
 */
func (this *CatalogController) GetCatalog() {
	userName := this.Ctx.GetCookie("username")
	user := models.GetUserByName(userName)
	catalogs := user.Catalog
	beego.Informational(catalogs)
	this.Ctx.WriteString("获取目录成功！")
}

//post request!
func (this *CatalogController) CreateCatalog() {
	beego.Informational("创建目录请求")
	frontCatalog := FrontCatalog{}
	if err := this.ParseForm(&frontCatalog); err != nil {
		//handle error
		beego.Informational("前端参数解析出错！")
	}
	token := this.GetSession(frontCatalog.CataUserName)
	if token != "-1" {
		catalog := models.GetCatalogByName(frontCatalog.CataName)
		if catalog != nil {
			beego.Informational("目录名已存在！")
			this.Ctx.WriteString("目录名已存在!")
		} else {
			newCatalog := models.Catalog{CatalogName:frontCatalog.CataName}
			success := models.InsertCatalog(&newCatalog)
			if success {
				beego.Informational("创建目录成功！")
				this.Ctx.WriteString("创建目录成功，目录名：" + frontCatalog.CataName)
			}
			beego.Informational("数据库插入目录失败！")
			this.Ctx.WriteString("数据库插入目录失败，目录名：" + frontCatalog.CataName)
		}
	}
	beego.Informational("请先登录！")
	this.Ctx.WriteString("创建失败，需要先登录帐号！")

}


//post request!
func (c *CatalogController) RenameCatalog() {

}


//post request!
func (c *CatalogController) DeleteCatalog() {

}


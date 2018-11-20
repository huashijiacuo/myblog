package routers

import (
	"myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/upload", &controllers.UploadController{})
    beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
    //beego.Router("/login2", &controllers.LoginController{}, "get:Get")
    beego.Router("/logout", &controllers.LogoutController{})
    beego.Router("/getCatalogs", &controllers.CatalogController{}, "get:GetCatalog")

}

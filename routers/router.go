package routers

import (
	"myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/upload", &controllers.UploadController{})
    beego.Router("/register", &controllers.RegisterController{})
}

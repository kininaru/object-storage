package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"object-storage/controllers"
)

func init() {
	beego.Router("/", &controllers.ApiController{}, "POST:Command")
	beego.Router("/*", &controllers.ApiController{}, "GET:GetFile")
}

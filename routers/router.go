package routers

import (
	"github.com/astaxie/beego"

	"github.com/mlgd/goHaxServer/controllers"
	"github.com/mlgd/goHaxServer/controllers/api"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.Router("/", &controllers.MainController{})
	beego.Router("/error/:numero", &controllers.ErrorController{}, "get:Get")
	beego.Router("/hax", &controllers.HaxController{}, "get:Get")
	beego.Router("/api/dns", &api.DnsController{}, "post:Post")
}

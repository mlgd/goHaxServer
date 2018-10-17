package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["dns_server"] = beego.AppConfig.DefaultBool("EnableDNSServer", false)
	c.Data["dns_filter"] = beego.AppConfig.DefaultBool("EnableDNSFilter", false)
	c.Data["version"] = beego.AppConfig.DefaultString("version", "1.0.0")
	c.Layout = "layout.tpl"
	c.TplName = "default.tpl"
}

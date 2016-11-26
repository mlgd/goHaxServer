package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["version"] = beego.AppConfig.DefaultString("version", "1.0.0")
	c.Layout = "layout.tpl"
	c.TplName = "default.tpl"
}

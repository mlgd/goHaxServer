package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Prepare() {
	c.Data["version"] = beego.AppConfig.DefaultString("version", "1.0.0")
}

func (c *ErrorController) Get() {
	switch c.Ctx.Input.Param(":numero") {
	case "401":
		c.Error401()
	case "403":
		c.Error403()
	case "500":
		c.Error500()
	case "503":
		c.Error503()
	default:
		c.Error404()
	}
}

func (c *ErrorController) Error401() {
	c.Data["error"] = 401
	c.Layout = "layout.tpl"
	c.TplName = "error.tpl"
}

func (c *ErrorController) Error403() {
	c.Data["error"] = 403
	c.Layout = "layout.tpl"
	c.TplName = "error.tpl"
}

func (c *ErrorController) Error404() {
	c.Data["error"] = 404
	c.Layout = "layout.tpl"
	c.TplName = "error.tpl"
}

func (c *ErrorController) Error500() {
	c.Data["error"] = 500
	c.Layout = "layout.tpl"
	c.TplName = "error.tpl"
}

func (c *ErrorController) Error503() {
	c.Data["error"] = 503
	c.Layout = "layout.tpl"
	c.TplName = "error.tpl"
}

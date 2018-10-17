package api

import (
	"github.com/astaxie/beego"
	"github.com/mlgd/goHaxServer/lib/dns"
)

type DnsController struct {
	beego.Controller
}

func (c *DnsController) Post() {
	if c.Ctx.Input.IsAjax() == false {
		c.Ctx.WriteString("")
		return
	}

	switch {
	case c.Input().Get("dns_server") != "":
		if err := beego.AppConfig.Set("enablednsserver", c.Input().Get("dns_server")); err != nil {
			beego.Critical(err)
		}
		if beego.AppConfig.DefaultBool("enablednsserver", false) {
			beego.Info("DNS server Running")
			dns.StartServer(beego.AppConfig.DefaultBool("enablednsfilter", false))
		} else {
			beego.Info("DNS server Stopping")
			dns.StopServer()
			beego.AppConfig.Set("enablednsfilter", "false")
		}

	case c.Input().Get("dns_filter") != "":
		if err := beego.AppConfig.Set("enablednsfilter", c.Input().Get("dns_filter")); err != nil {
			beego.Critical(err)
		}
		if beego.AppConfig.DefaultBool("enablednsfilter", false) {
			dns.EnableFilter()
		} else {
			dns.DisableFilter()
		}
	}
	if err := beego.AppConfig.SaveConfigFile("conf/app.conf"); err != nil {
		beego.Critical(err)
	}

	c.Ctx.WriteString("")
}

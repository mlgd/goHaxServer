package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/astaxie/beego"

	"github.com/mlgd/goHaxServer/lib/dns"
	_ "github.com/mlgd/goHaxServer/routers"
)

func main() {
	// Logs file
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		err = os.MkdirAll("log", 0755)
	}
	beego.SetLogger("file", `{"filename":"log/`+beego.AppConfig.String("appname")+`.log"}`)
	beego.SetLogFuncCall(true)

	// Signal for stop program
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c

		if beego.AppConfig.DefaultBool("enablednsserver", false) {
			beego.Info("DNS server Stopping")
			dns.StopServer()
		}

		beego.Info("http server Stopping")
		time.Sleep(50 * time.Millisecond)
		os.Exit(0)
	}()

	beego.AddAPPStartHook(appStart)

	beego.Run()
}

func appStart() error {
	dns.SetFilterAddresses(beego.AppConfig.DefaultStrings("dnsfilteraddresses", []string{}))
	if beego.AppConfig.DefaultBool("enablednsserver", false) {
		dns.StartServer(beego.AppConfig.DefaultBool("enablednsfilter", false))
		beego.Info("DNS server Running")
	}

	return nil
}

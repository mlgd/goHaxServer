package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/astaxie/beego"

	_ "github.com/mlgd/goHaxServer/routers"
)

const (
	VERSION = "1.0.0"
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
		beego.Info("http server Stopping")
		time.Sleep(50 * time.Millisecond)
		os.Exit(0)
	}()

	beego.Run()
}

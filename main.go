package main

import (
	_ "wx-session/docs"
	_ "wx-session/routers"

	"github.com/astaxie/beego"
	"wx-session/filter"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterUser)
	beego.Run()
}

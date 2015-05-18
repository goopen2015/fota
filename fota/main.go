package main

import (
	_ "fota/fota/docs"
	_ "fota/fota/routers"
	_ "fota/fota/models"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

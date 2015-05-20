package main

import (
	_ "fota/fota/docs"
	_ "fota/fota/models"
	_ "fota/fota/routers"

	_ "fota/fota/initial"
	"github.com/astaxie/beego"
)

func init() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}

}

func main() {

	beego.Run()
}

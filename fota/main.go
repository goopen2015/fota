package main

import (
	_ "fota/fota/docs"
	_ "fota/fota/routers"
	_ "fota/fota/models"

	"github.com/astaxie/beego"
	_ "fota/fota/initial"
)


func init(){
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}

}

func main() {

	beego.Run()
}

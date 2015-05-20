package initial

import (
	"github.com/astaxie/beego"
)

func init() {
	/*init the log*/
	err := beego.SetLogger("file", `{"filename":"/var/fota.log"}`)
	if err != nil {
		panic(err)
	}

	beego.SetLevel(beego.LevelDebug)
}

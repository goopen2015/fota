package models

import(
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)
func init(){
	//register the driver
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	//register the database
	orm.RegisterDataBase("default", "mysql", "root:root@/fota?charset=utf8")

	//register modules
	orm.RegisterModel(new(ProductFamily), new(ProductCU), new(Package),
		new(PFamilyPCU), new(PCUPackage))
	orm.RunSyncdb("default", false, true)

	beego.Debug("Models initiated")
}



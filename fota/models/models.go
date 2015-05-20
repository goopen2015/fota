package models

import(
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)
func init(){
	//register the driver
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	//register the database
	dbhost := beego.AppConfig.String("dhbost")
	dbport, err := beego.AppConfig.Int("dbport")
	if err != nil {
		dbport = 3306
	}
	dbuser := beego.AppConfig.String("dbuser")
	dbpwd := beego.AppConfig.String("dbpwd")
	dbname := beego.AppConfig.String("dbname")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbuser, dbpwd, dbhost, dbport, dbname)
	orm.RegisterDataBase("default", "mysql", connStr)

	//register modules
	orm.RegisterModel(new(ProductFamily), new(ProductCU), new(Package),
		new(PFamilyPCU), new(PCUPackage))
	orm.RunSyncdb("default", false, true)

	beego.Debug("Models initiated")
}



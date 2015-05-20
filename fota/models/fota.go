package models

import(
	"github.com/astaxie/beego/orm"
)


/*table struct*/

//product_family table
type ProductFamily struct{
	Id int `orm:"pk"`
	ProductFamilyName string `orm:"unique"`
	IsExisted bool
}

//product_cu table, ProductName+CUName is unique
type ProductCU struct{
	Id int `orm:"pk"`
	ProductName string
	CUName string
	IsExisted bool
}

//package table
type Package struct{
	Id int `orm:"pk"`
	PackageName string `orm:"unique"`
	IsExisted bool
}

//relation between product family table and product&cu table
type PFamilyPCU struct {
	Id int `orm:"pk"`
	ProductFamilyName string
	ProductName string
	CUName string
	IsExisted bool
}

//relation between product&cu table an package table
type PCUPackage struct{
	Id int `orm:"pk"`
	ProductName string
	CUName string
	PackageName string
	IsExisted bool
}

/*CRUD for product_family table*/
func InsertProductFamily(productFamilyName string){

	pf := new(ProductFamily)
	pf.ProductFamilyName = productFamilyName

	o := orm.NewOrm()

	_, err := o.Insert(pf)
	if err != nil {
		panic(err)
	}
}

func QueryProductFamilyByName(productFamilyName string) *ProductFamily{

	pf := new(ProductFamily)
	pf.ProductFamilyName = productFamilyName

	o := orm.NewOrm()

	err := o.Read(pf)
	if err != nil{
		panic(err)
	}else{
		return pf
	}
}



/*CRUD for product_cu table*/



/*CRUD for package table*/




package models

import(
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

/*Class for database operation*/
type CRUDOperator struct{

}

func (this *CRUDOperator) Insert(model interface{}){

	o := orm.NewOrm()

	_, err := o.Insert(model)
	if err != nil {
		panic(err)
	}
}

func (this *CRUDOperator) Query(model interface{}){

	o := orm.NewOrm()

	err := o.Read(model, "Name")
	if err != nil{
		panic(err)
	}
}

func (this *CRUDOperator) Update(model interface{}){
	o := orm.NewOrm()

	_, err := o.Update(model)
	if err != nil{
		panic(err)
	}
}

func (this *CRUDOperator) Delete(model interface{}){
	o := orm.NewOrm()

	_, err := o.Delete(model)
	if err != nil {
		panic(err)
	}
}

var opr CRUDOperator

//product_family table
type ProductFamily struct{
	Id int `orm:"pk"`
	Name string `orm:"unique"` //productfamily's name
	IsExisted bool
}

/*CRUD for product_family table*/
func (this *ProductFamily) InsertProductFamily(){

	opr.Insert(this)

	beego.Debug("Insert to ProductFamily table via productFamilyName is " + this.Name)

}

func (this *ProductFamily) QueryProductFamilyByName(){

	opr.Query(this)

	beego.Debug("Query from ProductFamily table where productFamilyName is " + this.Name)

}

func (this *ProductFamily) UpdateProductFamily(){

	opr.Update(this)

	beego.Debug("Update ProductFamily table via productFamilyName is " + this.Name)
}

func (this *ProductFamily) DelProductFamily(){

	opr.Delete(this)

	beego.Debug("Delete from ProductFamily table via productFamilyName is " + this.Name)
}


//product_cu table, ProductName+CUName is unique
type ProductCU struct{
	Id int `orm:"pk"`
	Name string //product+cu's name
	IsExisted bool
}
/*CRUD for product_cu table*/
func (this *ProductCU) InsertProductCU(productFamilyName string){
	opr.Insert(this)

	pFamilyPCU := &PFamilyPCU{Name:productFamilyName, ProductCUName:this.Name, IsExisted:true}
	opr.Insert(pFamilyPCU)

	beego.Debug("Insert to ProductCU table via ProductCUName is " + this.Name)
}

func (this *ProductCU) QueryProductCUByName(){

	opr.Query(this)

	beego.Debug("Query from ProductCU table where ProductCUName is " + this.Name)

}

func (this *ProductCU) UpdateProductCU(){

	opr.Update(this)

	beego.Debug("Update ProductCU table via ProductCUName is " + this.Name)
}

func (this *ProductCU) DelProductCU(){

	opr.Delete(this)

	beego.Debug("Delete from ProductCU table via ProductCUName is " + this.Name)
}

//package table
type Package struct{
	Id int `orm:"pk"`
	Name string `orm:"unique"`
	IsExisted bool
}
/*CRUD for package table*/
func (this *Package) InsertPackage(productCUName string){
	opr.Insert(this)

	pCUPackage := &PCUPackage{Name:productCUName, PackageName:this.Name, IsExisted:true}
	opr.Insert(pCUPackage)

	beego.Debug("Insert to Package table via PackageName is " + this.Name)
}

func (this *ProductCU) QueryPackageByName(){

	opr.Query(this)

	beego.Debug("Query from Package table where PackageName is " + this.Name)

}

func (this *ProductCU) UpdatePackage(){

	opr.Update(this)

	beego.Debug("Update Package table via PackageName is " + this.Name)
}

func (this *ProductCU) DelPackage(){

	opr.Delete(this)

	beego.Debug("Delete from Package table via PackageName is " + this.Name)
}


//relation between product family table and product&cu table
type PFamilyPCU struct {
	Id int `orm:"pk"`
	Name string //productfamily's name
	ProductCUName string
	IsExisted bool
}
/*CRUD for PFamilyPCU table*/
func (this *PFamilyPCU) InsertPFamilyPCU(){

	opr.Insert(this)

	beego.Debug("Insert to PFamilyPCU table via product familiy name is " + this.Name)

}

func (this *PFamilyPCU) QueryPFamilyPCU(){

	o := orm.NewOrm()

	err := o.Read(this, "Name", "ProductCUName")
	if err != nil{
		panic(err)
	}

	beego.Debug("Query from PFamilyPCU table where product familiy name is " + this.Name +
					"and productcu name is " + this.ProductCUName)
}

func (this *PFamilyPCU) UpdatePFamilyPCU(){

	opr.Update(this)

	beego.Debug("Update PFamilyPCU table via product familiy name is " + this.Name)
}

func (this *PFamilyPCU) DelPFamilyPCU(){

	opr.Delete(this)

	beego.Debug("Delete from PFamilyPCU table via product familiy name is " + this.Name)
}


//relation between product&cu table an package table
type PCUPackage struct{
	Id int `orm:"pk"`
	Name string //product+cu's name
	PackageName string
	IsExisted bool
}
/*CRUD for PCUPackage table*/
func (this *PCUPackage) InsertPCUPackage(){

	opr.Insert(this)

	beego.Debug("Insert to PCUPackage table via product and cu name is " + this.Name)

}

func (this *PCUPackage) QueryPCUPackage(){

	o := orm.NewOrm()

	err := o.Read(this, "Name", "PackageName")
	if err != nil{
		panic(err)
	}

	beego.Debug("Query from PCUPackage table where productcu name is " + this.Name +
	"and package name is " + this.PackageName)

}

func (this *PCUPackage) UpdatePCUPackage(){

	opr.Update(this)

	beego.Debug("Update PFamilyPCU table via product and cu name is " + this.Name)
}

func (this *PCUPackage) DelPCUPackage(){

	opr.Delete(this)

	beego.Debug("Delete from PFamilyPCU table via product and cu name is " + this.Name)
}




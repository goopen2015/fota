package models


/*table struct*/

//product family table
type ProductFamily struct{
	Id int `orm:"pk"`
	ProductFamilyName string `orm:"unique"`
	IsExisted bool
}

//product&cu table
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



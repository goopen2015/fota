package models


/*table struct*/

//product family table
type ProductFamily struct{
	Id string
	ProductFamilyName string
}

//product&cu table
type ProductCU struct{
	Id string
	ProductName string
	CUName string
}

//package table
type Package struct{
	Id string
	PackageName string
}

//relation between product family table and product&cu table
type PFamilyPCU struct {
	Id string
	ProductFamilyName string
	ProductName string
	CUName string
}

//relation between product&cu table an package table
type PCUPackage struct{
	Id string
	ProductName string
	CUName string
	PackageName string
}



package product_service

import (
	"staff-mall-center/models/dao"
	"time"
)

type Product struct {
	PID             int
	ProductName     string
	ProductRealname string

	CategoryID       int
	CategoryName     string
	CategoryRealname string

	ProductDesc   string
	ProductImg    string
	ProductCount  int
	ProductPrice  int
	ProductStatus int

	SupplierId       int
	SupplierName     string
	SupplierRealname string
	SupplierTel      int
}
type ArrayProduct []Product

func (plist *ArrayProduct) BuckRegister() ([]int, error) {
	nowTime := time.Now().Unix()
	product_list := make([]interface{}, len(*plist))
	for idx, val := range *plist {
		product_list[idx] = dao.Product{
			ProductName:      val.ProductName,
			ProductRealname:  val.ProductRealname,
			CategoryID:       val.CategoryID,
			CategoryName:     val.CategoryName,
			CategoryRealname: val.CategoryRealname,
			ProductCount:     val.ProductCount,
			ProductPrice:     val.ProductPrice,
			ProductImg:       val.ProductImg,
			ProductDesc:      val.ProductDesc,
			ProductStatus:    val.ProductStatus,

			SupplierId:       val.SupplierId,
			SupplierName:     val.SupplierName,
			SupplierRealname: val.SupplierRealname,
			SupplierTel:      val.SupplierTel,

			Model: dao.Model{
				CreatedOn: nowTime,
			},
		}
	}
	return dao.BuckUpsertProduct(product_list)
}

func (product *Product) GetInfo() error {
	_product, err := dao.GetProductItem(product.PID)

	if err != nil {
		return err
	}

	product.PID = _product.ID
	product.ProductName = _product.ProductName
	product.ProductRealname = _product.ProductRealname
	product.ProductPrice = _product.ProductPrice
	product.ProductStatus = _product.ProductStatus
	product.ProductDesc = _product.ProductDesc
	product.ProductImg = _product.ProductImg
	product.ProductCount = _product.ProductCount

	product.CategoryID = _product.CategoryID
	product.CategoryName = _product.CategoryName
	product.CategoryName = _product.CategoryName

	product.SupplierId = _product.SupplierId
	product.SupplierName = _product.SupplierName
	product.SupplierRealname = _product.SupplierRealname
	product.SupplierTel = _product.SupplierTel

	return err
}

func (product *Product) Occupy(count int) error {
	_product, err := dao.OccupyProductItem(product.PID, count)

	if err != nil {
		if err.Error() == "OverRemain" {
			product.PID = _product.ID
			product.ProductCount = _product.ProductCount
			return err
		}
		return err
	}

	product.PID = _product.ID
	product.ProductPrice = _product.ProductPrice

	return err
}

func (p *Product) Register() error {
	product, err := dao.RegisterProduct(
		p.ProductName,
		p.ProductRealname,
		p.CategoryID,
		p.CategoryName,
		p.CategoryRealname,
		p.ProductPrice,
		p.ProductCount,
		p.ProductImg,
		p.ProductStatus,
		p.ProductDesc,
		p.SupplierId,
		p.SupplierName,
		p.SupplierRealname,
		p.SupplierTel,
	)
	p.PID = product.ID
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) Find() error {
	product, err := dao.GetProductItem(
		p.PID,
	)
	if err != nil {
		return err
	}
	p.ProductName = product.ProductName
	p.ProductRealname = product.ProductRealname
	p.CategoryID = product.CategoryID
	p.CategoryName = product.CategoryName
	p.CategoryRealname = product.CategoryRealname
	p.ProductPrice = product.ProductPrice
	p.ProductCount = product.ProductCount
	p.ProductImg = product.ProductImg
	p.ProductStatus = product.ProductStatus
	p.ProductDesc = product.ProductDesc

	p.SupplierId = product.SupplierId
	p.SupplierName = product.SupplierName
	p.SupplierRealname = product.SupplierRealname
	p.SupplierTel = product.SupplierTel

	return nil
}

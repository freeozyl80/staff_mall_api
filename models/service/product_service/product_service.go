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
}
type ArrayProduct []Product

func (plist *ArrayProduct) BuckRegister() ([]int, error) {
	nowTime := time.Now().Unix()
	user_list := make([]interface{}, len(*plist))
	for idx, val := range *plist {
		user_list[idx] = dao.Product{
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
			Model: dao.Model{
				CreatedOn: nowTime,
			},
		}
	}
	return dao.BuckUpsertProduct(user_list)
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

	return err
}

func (product *Product) Occupy(count int) error {
	_product, err := dao.OccupyProductItem(product.PID, count)

	if err != nil {
		return err
	}

	product.PID = _product.ID
	product.ProductPrice = _product.ProductPrice

	return err
}

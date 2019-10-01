package dao

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Product struct {
	Model
	ProductName      string `gorm:"not null;unique" json:"product_name"`
	ProductRealname  string `json:"product_realname"`
	CategoryID       int    `json:"category_id`
	CategoryName     string `json:"category_name`
	CategoryRealname string `json:"category_realname`
	ProductCount     int    `json:"product_count"`
	ProductStatus    int    `gorm:"not null; default:1" json:"product_status"`
	ProductPrice     int    `gorm:"not null; default:1" json:"product_price"`
	ProductDesc      string `gorm:"not null; default:1" json:"product_desc"`
	ProductImg       string `gorm:"not null; default:'https://img.zcool.cn/community/01fd5958be2833a801219c77b15a63.jpg@2o.jpg'" json:"product_img"`

	SupplierId       int    `gorm:"not null; default:1" json:"supplier_id"`
	SupplierName     string `gorm:"not null;unique" json:"supplier_name"`
	SupplierRealname string `gorm:"not null" json:"supplier_realname"`
	SupplierTel      int    `json:"supplier_tel"`
}

func BuckUpsertProduct(objArr []interface{}) ([]int, error) {
	ids, err := BulkInsertOnDuplicateUpdate(db, objArr,
		"product_name = values(product_name),"+
			"product_realname = values(product_realname),"+
			"category_id = values(category_id),"+
			"category_name = values(category_name),"+
			"category_realname = values(category_realname),"+
			"product_count = values(product_count),"+
			"product_price = values(product_price),"+
			"product_img = values(product_img),"+
			"product_status = values(product_status),"+
			"product_desc = values(product_desc),"+
			"supplier_id = values(supplier_id),"+
			"supplier_name = values(supplier_name),"+
			"supplier_realname = values(supplier_realname),"+
			"supplier_tel = values(supplier_tel)")
	return ids, err
}

func GetProductList(pageIndex int, pageSize int, maps interface{}) ([]*Product, error) {
	var productList []*Product
	err := db.Where(maps).Offset(pageIndex).Limit(pageSize).Find(&productList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return productList, nil
}

func GetProductItem(id int) (Product, error) {

	var product Product

	err := db.Where("id = ?", id).First(&product).Error

	// 存在
	if product.ID > 0 && err == nil {
		return product, nil
	} else {
		return product, errors.New("can not find")
	}
}
func OccupyProductItem(id int, count int) (Product, error) {
	var product Product

	err := db.Where("id = ?", id).First(&product).Error

	// 存在
	if product.ID > 0 && err == nil && (product.ProductCount-count) >= 0 {

		var leftCount = product.ProductCount - count
		err = db.Model(&Product{}).Where("id = ?", id).Update("ProductCount", leftCount).Error
		if err != nil {
			return product, errors.New("occpuy faild")
		}
		return product, nil
	} else if product.ProductCount-count < 0 {
		return product, errors.New("OverRemain")
	} else {
		return product, errors.New("can not find")
	}
}

func RegisterProduct(
	productName string,
	productRealname string,
	categoryID int,
	categoryName string,
	categoryRealname string,
	productPrice int,
	productCount int,
	productImg string,
	productStatus int,
	productDesc string,

	supplier_id int,
	supplier_name string,
	supplier_realname string,
	supplier_tel int,
) (Product, error) {
	var product = Product{
		ProductName:      productName,
		ProductRealname:  productRealname,
		CategoryID:       categoryID,
		CategoryName:     categoryName,
		CategoryRealname: categoryRealname,
		ProductPrice:     productPrice,
		ProductCount:     productCount,
		ProductImg:       productImg,
		ProductStatus:    productStatus,
		ProductDesc:      productDesc,

		SupplierId:       supplier_id,
		SupplierName:     supplier_name,
		SupplierRealname: supplier_realname,
		SupplierTel:      supplier_tel,
	}

	if err := db.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil

}
func UpdateProduct(id int, data interface{}) error {
	if err := db.Model(&Product{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

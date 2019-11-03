package dao

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Supplier struct {
	Model

	SupplierName     string `gorm:"not null;unique" json:"supplier_name"`
	SupplierRealname string `gorm:"not null" json:"supplier_realname"`
	SupplierTel      string `json:"supplier_tel"`
}

func BuckUpsertSupplier(objArr []interface{}) ([]int, error) {
	fmt.Printf("%+v\n", objArr)
	fmt.Println("--------------------------")
	ids, err := BulkInsertOnDuplicateUpdate(db, objArr,
		"supplier_name = values(supplier_name),"+
			"supplier_tel = values(supplier_tel),"+
			"supplier_realname = values(supplier_realname)")
	fmt.Println(err)

	return ids, err
}

func FindSupplierById(sname string) (Supplier, error) {

	var supplier Supplier

	err := db.Where(Supplier{SupplierName: sname}).First(&supplier).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		return supplier, errors.New("not found supplier")
	}

	return supplier, err
}
func GetSupplierList(pageIndex int, pageSize int, maps interface{}) ([]*Supplier, error) {
	var supplierList []*Supplier
	err := db.Where(maps).Offset(pageIndex).Limit(pageSize).Find(&supplierList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return supplierList, nil
}

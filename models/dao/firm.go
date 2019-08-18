package dao

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Firm struct {
	Model
	//	企业相关信息
	Firmname     string `gorm:"not null;unique" json:"firm_name"`
	FirmRealname string `json:"firm_realname"`
	// 企业其他信息

	Balance int `gorm:"not null; default:0" json:"balance"`

	// 关联商品类别
	CategoryGroup string `json:"category_group"`
	// 关联商品
	ProductGroup string `json:"product_group"`

	// 状态，可以被禁用掉 1 是正常 2是被禁用掉
	FirmStatus int `gorm:"not null; default:1" json:"firm_status"`
}

func BuckUpsertFirm(objArr []interface{}) ([]int, error) {
	ids, err := BulkInsertOnDuplicateUpdate(db, objArr,
		"firmname = values(Firmname), firm_realname = values(FirmRealname), balance = values(Balance), firmtatus = values(FirmStatus)")
	return ids, err
}

func CreateFirm(firm_status, balance int, firmname, firm_realname, category_group, product_group string) error {
	var firm = Firm{
		Firmname:      firmname,
		FirmRealname:  firm_realname,
		Balance:       balance,
		CategoryGroup: category_group,
		ProductGroup:  product_group,
		FirmStatus:    firm_status,
	}

	if err := db.Create(&firm).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GetFirmList(pageIndex int, pageSize int, maps interface{}) ([]*Firm, error) {
	var firmList []*Firm
	err := db.Where(maps).Offset(pageIndex).Limit(pageSize).Find(&firmList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return firmList, nil
}

func UpdateFirm(id int, data interface{}) error {
	if err := db.Model(&Firm{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func GetFirmItem(id int) (Firm, error) {
	var firm Firm

	err := db.Where("id = ?", id).First(&firm).Error

	// 存在
	if firm.ID > 0 && err == nil {
		return firm, nil
	} else {
		return firm, errors.New("can not find")
	}
}

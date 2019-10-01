package dao

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Category struct {
	Model
	CategoryName     string `gorm:"not null;unique" json:"category_name"`
	CategoryRealname string `json:"category_realname"`
	CategoryStatus   int    `gorm:"not null; default:1" json:"category_status"`
	CategoryDesc     string `gorm:"not null; default:'默认类别'" json:"category_desc"`
}

func BuckUpsertCategory(objArr []interface{}) ([]int, error) {

	fmt.Printf("%+v\n", objArr)
	fmt.Println("--------------------------")
	ids, err := BulkInsertOnDuplicateUpdate(db, objArr,
		"category_name = values(category_name),"+
			"category_realname = values(category_realname),"+
			"category_status = values(category_status),"+
			"category_desc = values(category_desc)")
	fmt.Println(err)

	return ids, err
}

func FindCategoryId(cname string) (Category, error) {

	var category Category

	err := db.Where(Category{CategoryName: cname}).First(&category).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		return category, errors.New("not found category")
	}

	return category, err
}

func GetCategoryList(pageIndex int, pageSize int, maps interface{}) ([]*Category, error) {
	var categoryList []*Category
	err := db.Where(maps).Offset(pageIndex).Limit(pageSize).Find(&categoryList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return categoryList, nil
}

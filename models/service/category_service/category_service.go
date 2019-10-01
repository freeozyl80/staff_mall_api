package category_service

import (
	"staff-mall-center/models/dao"
	"time"
)

type Category struct {
	CID              int
	CategoryName     string
	CategoryRealname string
	CategoryDesc     string
	CategoryStatus   int
}
type ArrayCategory []Category

func (clist *ArrayCategory) BuckRegister() ([]int, error) {
	nowTime := time.Now().Unix()
	category_list := make([]interface{}, len(*clist))
	for idx, val := range *clist {
		category_list[idx] = dao.Category{
			CategoryName:     val.CategoryName,
			CategoryRealname: val.CategoryRealname,
			CategoryDesc:     val.CategoryDesc,
			CategoryStatus:   val.CategoryStatus,
			Model: dao.Model{
				CreatedOn: nowTime,
			},
		}
	}
	return dao.BuckUpsertCategory(category_list)
}

func (c *Category) FindCategory() error {
	category, err := dao.FindCategoryId(c.CategoryName)

	if err != nil {
		return err
	}

	c.CID = category.ID
	return nil
}

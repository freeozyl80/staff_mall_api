package firm_service

import (
	"fmt"
	"staff-mall-center/models/dao"
	"time"
)

type Firm struct {
	Fid          int
	UID          int
	Username     string
	Firmname     string
	FirmRealname string

	Balance       int
	CategoryGroup string
	ProductGroup  string
	FirmStatus    int
}
type ArrayFirm []Firm

func (firmlist *ArrayFirm) BuckRegister() ([]int, error) {
	nowTime := time.Now().Unix()
	firm_list := make([]interface{}, len(*firmlist))
	for idx, val := range *firmlist {
		firm_list[idx] = dao.Firm{
			Username:      val.Username,
			UID:           val.UID,
			Balance:       val.Balance,
			FirmRealname:  val.FirmRealname,
			Firmname:      val.Firmname,
			CategoryGroup: val.CategoryGroup,
			ProductGroup:  val.ProductGroup,
			FirmStatus:    val.FirmStatus,
			Model: dao.Model{
				CreatedOn: nowTime,
			},
		}
	}
	return dao.BuckUpsertFirm(firm_list)
}

func (firm *Firm) CreateFirm() (bool, error) {
	err := dao.CreateFirm(firm.UID, firm.FirmStatus, firm.Balance, firm.Firmname, firm.FirmRealname, firm.Username, firm.CategoryGroup, firm.CategoryGroup)
	if err != nil {
		fmt.Println("internal Error")
		return false, err
	}
	return true, nil
}

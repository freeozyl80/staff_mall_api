package firm_service

import (
	"staff-mall-center/models/dao"
	"time"
)

type Firm struct {
	Fid          int
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
	err := dao.CreateFirm(firm.FirmStatus, firm.Balance, firm.Firmname, firm.FirmRealname, firm.CategoryGroup, firm.CategoryGroup)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (firm *Firm) FindFirm() error {
	f, err := dao.GetFirmItem(firm.Fid)
	if err != nil {
		return err
	}
	firm.Firmname = f.Firmname
	firm.FirmRealname = f.FirmRealname
	firm.CategoryGroup = f.CategoryGroup
	firm.Balance = f.Balance
	firm.ProductGroup = f.ProductGroup
	return err
}

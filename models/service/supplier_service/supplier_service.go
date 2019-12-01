package supplier_service

import (
	"staff-mall-center/models/dao"
	"time"
)

type Supplier struct {
	SupplierId       int
	SupplierName     string
	SupplierRealname string
	SupplierTel      int
}

type ArraySupplier []Supplier

func (clist *ArraySupplier) BuckRegister() ([]int, error) {
	nowTime := time.Now().Unix()
	supplier_list := make([]interface{}, len(*clist))
	for idx, val := range *clist {
		supplier_list[idx] = dao.Supplier{
			SupplierName:     val.SupplierName,
			SupplierRealname: val.SupplierRealname,
			Model: dao.Model{
				CreatedOn: nowTime,
			},
		}
	}
	return dao.BuckUpsertSupplier(supplier_list)
}

func (s *Supplier) FindSupplier() error {
	supplier, err := dao.FindSupplierById(s.SupplierName)

	if err != nil {
		return err
	}

	s.SupplierId = supplier.ID
	return nil
}

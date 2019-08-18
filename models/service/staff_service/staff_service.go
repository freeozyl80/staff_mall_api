package staff_service

import (
	"staff-mall-center/models/dao"
	"time"
)

type Staff struct {
	SID      int
	UID      int
	Username string

	Realname     string
	FID          int
	Firmname     string
	FirmRealname string

	Gender   int
	Addresss string
	Avartar  string

	Tel      int
	Birthday int
	Coin     int
}
type ArrayStaff []Staff

func (plist *ArrayStaff) BuckRegister() ([]int, error) {
	nowTime := time.Now().Unix()
	user_list := make([]interface{}, len(*plist))
	for idx, val := range *plist {
		user_list[idx] = dao.Staff{
			UID:          val.UID,
			Username:     val.Username,
			Realname:     val.Realname,
			Tel:          val.Tel,
			Birthday:     val.Birthday,
			Coin:         val.Coin,
			Firmname:     val.Firmname,
			FirmRealname: val.FirmRealname,
			Gender:       val.Gender,
			Addresss:     val.Addresss,
			Avartar:      val.Avartar,
			Model: dao.Model{
				CreatedOn: nowTime,
			},
		}
	}
	return dao.BuckUpsertStaff(user_list)
}

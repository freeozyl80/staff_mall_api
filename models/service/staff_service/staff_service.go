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
	Fid          int
	Firmname     string
	FirmRealname string

	Gender      int
	UserAddress string
	UserAvatar  string

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
			Fid:          val.Fid,
			Firmname:     val.Firmname,
			FirmRealname: val.FirmRealname,
			Gender:       val.Gender,
			UserAddress:  val.UserAddress,
			UserAvatar:   val.UserAvatar,
			Model: dao.Model{
				CreatedOn: nowTime,
			},
		}
	}
	return dao.BuckUpsertStaff(user_list)
}

func (staff *Staff) GetInfo() error {

	_staff, err := dao.GetStaffItem(staff.UID)

	if err != nil {
		return err
	}

	staff.UID = _staff.UID
	staff.Username = _staff.Username
	staff.Realname = _staff.Realname
	staff.Tel = _staff.Tel
	staff.Birthday = _staff.Birthday
	staff.Coin = _staff.Coin
	staff.Fid = _staff.Fid
	staff.Firmname = _staff.Firmname
	staff.FirmRealname = _staff.FirmRealname
	staff.Gender = _staff.Gender
	staff.UserAddress = _staff.UserAddress
	staff.UserAvatar = _staff.UserAvatar
	return err
}

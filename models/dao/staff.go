package dao

import (
	"errors"
)

type Staff struct {
	Model

	UID      int    `gorm:"not null;unique" json:"uid"`
	Username string `gorm:"not null;unique" json:"username"`
	Realname string `gorm:"not null;" json:"realname"`

	Fid          int    `json:"fid"`
	Firmname     string `json:"firmname"`
	FirmRealname string `json:"firm_realname"`

	Gender      int    `gorm:"not null;default:1"json:"gender"`
	Tel         string `json:"tel"`
	Coin        int    `gorm:"not null;default:0" json:"coin"`
	Birthday    int    `json:"birthday"`
	UserAddress string `json:"user_address"`
	UserAvatar  string `gorm:"not null;default'//c-ssl.duitang.com/uploads/item/201704/27/20170427155254_Kctx8.thumb.700_0.jpeg'" json:"user_avatar"`
}

func BuckUpsertStaff(objArr []interface{}) ([]int, error) {
	ids, err := BulkInsertOnDuplicateUpdate(db, objArr,
		"uid = values(uid),"+
			"username = values(username),"+
			"realname = values(realname),"+
			"fid = values(fid),"+
			"firmname = values(firmname),"+
			"firm_realname = values(firm_realname),"+
			"gender = values(gender),"+
			"tel = values(tel),"+
			"coin = values(coin),"+
			"birthday = values(birthday),"+
			"user_address = values(user_address),"+
			"user_avatar = values(user_avatar)")
	return ids, err
}

func GetStaffItem(id int) (Staff, error) {
	var staff Staff

	err := db.Where("uid = ?", id).First(&staff).Error

	// 存在
	if staff.ID > 0 && err == nil {
		return staff, nil
	} else {
		return staff, errors.New("can not find")
	}
}

func UpdateStaffItem(uid int, data interface{}) error {

	if err := db.Model(&Staff{}).Where(Staff{UID: uid}).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func UpdateStaffCoin(uid int, fid int, coin int, basecoin int, firmcoin int) error {

	tx := db.Begin()

	if err := tx.Model(&Staff{}).Where(Staff{UID: uid, Fid: fid}).Update("coin", coin+basecoin).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&Firm{}).Where(Firm{Model: Model{
		ID: fid,
	}}).Update("balance", firmcoin-coin).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

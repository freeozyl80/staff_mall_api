package dao

import (
	"errors"
	u "staff-mall-center/pkg/user"
	"strconv"

	"staff-mall-center/pkg/setting"

	"github.com/jinzhu/gorm"
)

type User struct {
	Model

	Username string `gorm:"not null;unique" json:"username"`
	Password string `json:"password"`
	Usertype int    `json:"usertype"`
	Realname string `json:"realname"`
	Salt1    string `json:"salt1"`
	Salt2    string `json:"salt2"`
}

func GetUserItem(username string, usertype int) (User, error) {

	var user User

	err := db.Select("id").Where(User{Username: username, Usertype: usertype}).First(&user).Error

	// 存在
	if user.ID > 0 {
		return user, nil
	}
	// 有报错，且报错不是是没有找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}

	return user, nil
}

func CheckUser(username, password string, usertype int) (int, error) {
	u.CryptoHandler(&password)

	var user User
	var err error

	if usertype != 0 {
		err = db.Model(&User{}).Where(User{Username: username, Password: password, Usertype: usertype}).First(&user).Error
		if user.ID > 0 {
			return user.ID, nil
		}
	} else {
		err = db.Model(&User{}).Where(User{Username: username, Password: password}).First(&user).Error

		if user.ID > 0 && (user.Usertype == 2 || user.Usertype == 3) {
			return user.ID, nil
		}
	}
	// 存在

	// 有报错，且报错不是是没有找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return 0, errors.New("can't find res")
}

func CheckUid(uid int, password string) (int, error) {
	u.CryptoHandler(&password)

	var user User
	err := db.Select("id").Where(User{Model: Model{
		ID: uid,
	}, Password: password}).First(&user).Error

	// 存在
	if user.ID > 0 {
		return user.ID, nil
	}
	// 有报错，且报错不是是没有找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return 0, errors.New("can't find res")
}

func RegisterUser(username, password string, usertype int, realname, auth1 string) error {
	tx := db.Begin()

	u.CryptoHandler(&password)
	var user = User{
		Username: username,
		Password: password,
		Usertype: usertype,
		Realname: realname,
		Salt1:    setting.CryptoSetting.Seed1,
		Salt2:    setting.CryptoSetting.Seed2}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&Auth{Username: username,
		Auth1: auth1, Auth2: "1", Auth3: "1", UID: user.Model.ID}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if auth1 == "0" {

		tx.Commit()

		return nil
	}

	fid, _ := strconv.Atoi(auth1)

	firmItem, err := GetFirmItem(
		fid,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Staff{
		UID:          user.Model.ID,
		Username:     username,
		Realname:     realname,
		Fid:          firmItem.Model.ID,
		Firmname:     firmItem.Firmname,
		FirmRealname: firmItem.FirmRealname,
		Coin:         0,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func BuckUpsertUser(objArr []interface{}) ([]int, error) {
	ids, err := BulkInsertOnDuplicateUpdate(db, objArr,
		"username = values(Username), password = values(Password), usertype = values(Usertype), realname = values(Realname),  salt1 = values(Salt1), salt2 = values(Salt2)")
	return ids, err
}

func GetAccountList(total *int, pageIndex int, pageSize int, maps interface{}) ([]*User, error) {
	var userlist []*User
	err := db.Model(&User{}).Where(maps).Count(total).Offset(pageIndex).Limit(pageSize).Find(&userlist).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return userlist, nil
}

func UpdateUser(id int, data interface{}) error {
	if err := db.Model(&User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

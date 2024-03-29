package account_service

import (
	"fmt"
	"staff-mall-center/models/dao"
	"staff-mall-center/pkg/setting"
	"time"
)

type User struct {
	UID      int
	Username string
	Password string
	Realname string
	Usertype int
}

type FirmUser struct {
	UID      int
	Usertype int
	Username string
	Password string
	Realname string
	Auth1    string
}

type ArrayUser []User

func (u *FirmUser) AccountRegister() error {
	err := dao.RegisterUser(u.Username, u.Password, u.Usertype, u.Realname, u.Auth1)
	return err
}

func (u *User) Check() (bool, error) {
	uid, err := dao.CheckUser(u.Username, u.Password, u.Usertype)

	if uid == 0 || err != nil {
		fmt.Println("登录失败")
		return false, nil
	}

	if err != nil {
		fmt.Println("internal Error")
		return true, err
	}

	u.UID = uid
	return true, nil
}
func (u *User) CheckPwd() (bool, error) {
	uid, err := dao.CheckUid(u.UID, u.Password)

	if uid == 0 || err != nil {
		fmt.Println("登录失败")
		return false, nil
	}

	if err != nil {
		fmt.Println("internal Error")
		return true, err
	}

	u.UID = uid
	return true, nil
}
func (ulist *ArrayUser) BuckRegister() ([]int, error) {
	nowTime := time.Now().Unix()
	user_list := make([]interface{}, len(*ulist))
	for idx, val := range *ulist {
		user_list[idx] = dao.User{
			Username: val.Username,
			Password: val.Password,
			Realname: val.Realname,
			Usertype: val.Usertype,
			Salt1:    setting.CryptoSetting.Seed1,
			Salt2:    setting.CryptoSetting.Seed2,
			Model: dao.Model{
				CreatedOn: nowTime,
			},
		}
	}
	return dao.BuckUpsertUser(user_list)
}

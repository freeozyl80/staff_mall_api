package account_service

import (
	"fmt"
	"staff-mall-center/models/dao"
	"staff-mall-center/pkg/setting"
	"time"
)

type User struct {
	Username string
	Password string
	Realname string
	Usertype int
}
type ArrayUser []User

func (u *User) Register() (bool, error) {
	// 注册用户就看姓名 和 usertype 相同不相同
	isExist, err := dao.ExistUser(u.Username, u.Usertype)

	if !isExist && err == nil {
		fmt.Println("验证了 不存在，开始注册")
		err = dao.RegisterUser(u.Username, u.Password, u.Usertype, u.Realname)
	}
	if isExist {
		fmt.Println("当前用户已存在")
		return false, err
	}
	if err != nil {
		fmt.Println("internal Error")
		return false, err
	}
	return true, nil
}

func (u *User) Check() (bool, error) {
	isExist, err := dao.CheckUser(u.Username, u.Password, u.Usertype)

	if !isExist && err == nil {
		fmt.Println("登录失败")
		return false, nil
	}

	if err != nil {
		fmt.Println("internal Error")
		return true, err
	}

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

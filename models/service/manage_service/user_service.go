package manage_service

import (
	"fmt"
	"staff-mall-center/models/dao"
)

type User struct {
	Username string
	Password string
	Realname string
	Usertype int
}

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

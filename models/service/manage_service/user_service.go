package manage_service

import (
	"fmt"
	"staff-mall-center/models/dao"
)

type User struct {
	Username string
	Password string
	Usertype int
}

func (u *User) Register() (bool, error) {
	isExist, err := dao.ExistUser(u.Username, u.Password)
	if !isExist && err == nil {
		fmt.Println("验证了 不存在，开始注册")
		err = dao.RegisterUser(u.Username, u.Password, u.Usertype)
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

package auth_service

import (
	"fmt"
	"staff-mall-center/models/dao"
	"time"
)

type Auth struct {
	AID      int
	UID      int
	Username string
	Auth1    string
	Auth2    string
	Auth3    string
}
type ArrayAuth []Auth

func (alist *ArrayAuth) BuckRegister() ([]int, error) {
	nowTime := time.Now().Unix()
	auth_list := make([]interface{}, len(*alist))
	for idx, val := range *alist {
		auth_list[idx] = dao.Auth{
			Username: val.Username,
			UID:      val.UID,
			Auth1:    val.Auth1,
			Auth2:    val.Auth2,
			Auth3:    val.Auth3,
			Model: dao.Model{
				CreatedOn: nowTime,
			},
		}
	}
	return dao.BuckUpsertAuth(auth_list)
}

func (auth *Auth) Check() (bool, error) {
	au, err := dao.CheckAuth(auth.Username, auth.UID)

	if err != nil {
		fmt.Println("权限获取失败")
		return false, nil
	}

	auth.Auth1 = au.Auth1
	auth.Auth2 = au.Auth2
	auth.Auth3 = au.Auth3
	auth.AID = au.ID

	return true, nil
}

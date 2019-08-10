package auth_service

import (
	"staff-mall-center/models/dao"
	"time"
)

type Auth struct {
	UID      int
	Username string
	Auth1    int
	Auth2    int
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
			Model: dao.Model{
				CreatedOn: nowTime,
			},
		}
	}
	return dao.BuckUpsertAuth(auth_list)
}

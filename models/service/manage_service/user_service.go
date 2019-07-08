package manage_service

import (
	"staff-mall-center/models/dao"
)

type User struct {
	Username string
	Password string
	Usertype int
}

func (u *User) Register() (bool, error) {
	return dao.RegisterUser(u.Username, u.Password, u.Usertype)
}

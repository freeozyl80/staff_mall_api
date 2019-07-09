package dao

import "github.com/jinzhu/gorm"

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	Usertype int    `json:"usertype"`
}

func ExistUser(username, password string) (bool, error) {
	var user User

	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error

	// 存在
	if user.ID > 0 {
		return true, nil
	}

	// 有报错，且报错不是是没有找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return false, nil
}

func RegisterUser(username, password string, usertype int) (error) {
	var user = User{Username: username, Password: password, Usertype: usertype}
	if err := db.Create(&user) {
		return err
	}
	return nil
}

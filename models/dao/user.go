package dao

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	Usertype int    `json:"usertype"`
}

func RegisterUser(username, password string, usertype int) (bool, error) {
	// var user User
	// // 判断有无
	// err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error

	// fmt.Printf("注册用户名是 %v，密码为：%v usertype是 %v \n", username, password, usertype)
	// // 已经存在
	// if user.ID > 0 {
	// 	return true, nil
	// }
	var user = User{Username: username, Password: password, Usertype: usertype}

	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return false, err
	// }

	db.Create(&user)

	// if err != nil {
	// 	return false, err
	// }
	return false, nil
}

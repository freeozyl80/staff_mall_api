package dao

import (
	"fmt"
	u "staff-mall-center/pkg/user"

	"staff-mall-center/pkg/setting"

	"github.com/jinzhu/gorm"
)

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	Usertype int    `json:"usertype"`
	Realname string `json:"realname"`
	Salt1    string `json:"salt1"`
	Salt2    string `json:"salt2"`
}

func ExistUser(username string, usertype int) (bool, error) {

	var user User

	err := db.Select("id").Where(User{Username: username, Usertype: usertype}).First(&user).Error

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

func CheckUser(username, password string, usertype int) (bool, error) {
	u.CryptoHandler(&password)

	var user User

	err := db.Select("id").Where(User{Username: username, Password: password, Usertype: usertype}).First(&user).Error

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

func RegisterUser(username, password string, usertype int, realname string) error {

	u.CryptoHandler(&password)

	var user = User{Username: username, Password: password, Usertype: usertype, Realname: realname, Salt1: setting.CryptoSetting.Seed1, Salt2: setting.CryptoSetting.Seed2}

	fmt.Println(user)

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// func BulkInsert(unsavedRows []*ExampleRowStruct) error {
// 	valueStrings := make([]string, 0, len(unsavedRows))
// 	valueArgs := make([]interface{}, 0, len(unsavedRows) * 3)
// 	for _, post := range unsavedRows {
// 			valueStrings = append(valueStrings, "(?, ?, ?)")
// 			valueArgs = append(valueArgs, post.Column1)
// 			valueArgs = append(valueArgs, post.Column2)
// 			valueArgs = append(valueArgs, post.Column3)
// 	}
// 	stmt := fmt.Sprintf("INSERT INTO my_sample_table (column1, column2, column3) VALUES %s", strings.Join(valueStrings, ","))
// 	_, err := db.Exec(stmt, valueArgs...)
// 	return err
// }

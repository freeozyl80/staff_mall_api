package dao

import (
	"errors"
	"staff-mall-center/pkg/setting"

	"github.com/jinzhu/gorm"
)

type Auth struct {
	Model

	UID      int    `gorm:"not null;unique" json:"uid"`
	Username string `json:"username"`
	Auth1    string `gorm:"default:'1'" json:"auth1"`
	Auth2    string `gorm:"default:'1'" json:"auth2"`
	Auth3    string `gorm:"default:'1'" json:"auth3"`
}

func (Auth) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "auth"
}

func CheckAuth(username string, uid int) (Auth, error) {
	var auth Auth

	err := db.Where(Auth{UID: uid, Username: username}).First(&auth).Error

	// 存在
	if auth.ID > 0 {
		return auth, nil
	}

	// 有报错，且报错不是是没有找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, errors.New("can not find res")
}

func BuckUpsertAuth(objArr []interface{}) ([]int, error) {
	ids, err := BulkInsertOnDuplicateUpdate(db, objArr,
		"username = values(Username), uid = values(UID), auth1 = values(Auth1), auth2 = values(Auth2), auth3 = values(Auth3)")
	return ids, err
}

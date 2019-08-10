package dao

type Auth struct {
	Model

	UID      int    `json:"uid"`
	Username string `json:"username"`
	Auth1    int    `gorm:"default:1" json:"auth1"`
	Auth2    int    `gorm:"default:1" json:"auth2"`
}

func CheckAuth(username, password string) bool {
	// var auth Auth
	// db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	// if auth.ID > 0 {
	// 	return true
	// }

	return false
}

func BuckUpsertAuth(objArr []interface{}) ([]int, error) {
	ids, err := BulkInsertOnDuplicateUpdate(db, objArr,
		"username = values(Username), uid = values(UID), auth1 = values(Auth1), auth2 = values(Auth2)")
	return ids, err
}

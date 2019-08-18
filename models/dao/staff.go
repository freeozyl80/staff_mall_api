package dao

type Staff struct {
	Model

	UID      int    `gorm:"not null;unique" json:"uid"`
	Username string `gorm:"not null;unique" json:"username"`
	Realname string `gorm:"not null;unique" json:"realname"`

	Fid          int    `json:"fid"`
	Firmname     string `json:"firmname"`
	FirmRealname string `json:"firm_realname"`

	Gender   int    `gorm:"not null;default:1"json:"gender"`
	Tel      int    `json:"tel"`
	Coin     int    `gorm:"not null;default:0" json:"coin"`
	Birthday int    `json:"birthday"`
	Addresss string `json:"address"`
	Avartar  string `gorm:"not null;default'https://c-ssl.duitang.com/uploads/item/201704/27/20170427155254_Kctx8.thumb.700_0.jpeg'" json:"avatar"`
}

func BuckUpsertStaff(objArr []interface{}) ([]int, error) {
	ids, err := BulkInsertOnDuplicateUpdate(db, objArr,
		"uid = values(uid),"+
			"username = values(username),"+
			"realname = values(realname),"+
			"fid = values(fid),"+
			"firmname = values(firmname),"+
			"firm_realname = values(firm_realname),"+
			"gender = values(gender),"+
			"tel = values(tel),"+
			"coin = values(coin),"+
			"birthday = values(birthday),"+
			"address = values(Addresss),"+
			"avatar = values(Avartar)")
	return ids, err
}

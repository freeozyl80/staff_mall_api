package dao

import "staff-mall-center/pkg/setting"

type Firm struct {
	Model
	//	企业相关信息
	Fid          int    `gorm:"not null;unique" json:"fid"`
	Firmname     string `json:"firm_name"`
	FirmRealname string `json:"firm_realname"`

	// 企业负责人相关新
	UID      int    `gorm:"not null;unique" json:"uid"`
	Username string `json:"username"`

	// 企业其他信息

	Balance int `gorm:"not null; default:0" json:"balance"`

	// 关联商品类别
	CategoryGroup string `json:"category_group"`
	// 关联商品
	ProductGroup string `json:"product_group"`

	// 状态，可以被禁用掉 1 是正常 2是被禁用掉
	FirmStatus int `gorm:"not null; default:1" json:"firm_status"`
}

func (Firm) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "firm"
}

func BuckUpsertFirm(objArr []interface{}) ([]int, error) {
	ids, err := BulkInsertOnDuplicateUpdate(db, objArr,
		"firmname = values(Firmname), firm_realname = values(FirmRealname), uid = values(UID), username = values(Username), balance = values(Balance), firmtatus = values(FirmStatus)")
	return ids, err
}

func CreateFirm(uid, firm_status, balance int, firmname, firm_realname, username, category_group, product_group string) error {
	var firm = Firm{
		UID:           uid,
		Username:      username,
		Firmname:      firmname,
		FirmRealname:  firm_realname,
		Balance:       balance,
		CategoryGroup: category_group,
		ProductGroup:  product_group,
		FirmStatus:    firm_status,
	}

	if err := db.Create(&firm).Error; err != nil {
		return err
	}

	return nil
}

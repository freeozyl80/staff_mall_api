package bootstrap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/account_service"
	"staff-mall-center/models/service/auth_service"
	u "staff-mall-center/pkg/user"
)

var jsonstruct struct {
	Accountlist []struct {
		Username string
		Password string
		Realname string
		Usertype int
	}
}

func Setup() {
	err1 := importSuperManager()
	if err1 != nil {
		log.Println("导入 超管 账号失败")
	}
}

func importSuperManager() error {
	if istablish := dao.IsEstablish(); istablish {
		return nil
	}
	fmt.Println("数据库 建表完成，开始导入数据")

	var accoutList = account_service.ArrayUser{}
	var authList = auth_service.ArrayAuth{}

	absPath, _ := filepath.Abs("./account.json")
	data, err := ioutil.ReadFile(absPath)
	err = json.Unmarshal(data, &jsonstruct)

	for _, val := range jsonstruct.Accountlist {
		User := account_service.User{
			Usertype: val.Usertype,
			Username: val.Username,
			Realname: val.Realname,
			Password: val.Password,
		}
		u.CryptoHandler(&User.Password)
		accoutList = append(accoutList, User)

		Auth := auth_service.Auth{
			Username: val.Username,
			Auth1:    "1",
			Auth2:    "1",
			Auth3:    "1",
		}
		authList = append(authList, Auth)
	}
	ids, err := accoutList.BuckRegister()

	if err != nil {
		return err
	}

	for idx, val := range ids {
		authList[idx].UID = val
	}
	_, err = authList.BuckRegister()

	if err != nil {
		return err
	}

	return err
}

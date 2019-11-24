package user

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/account_service"
	"staff-mall-center/models/service/auth_service"
	"staff-mall-center/models/service/firm_service"
	"staff-mall-center/models/service/staff_service"
	util "staff-mall-center/pkg/auth"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"staff-mall-center/pkg/setting"
	"staff-mall-center/pkg/tools"
	u "staff-mall-center/pkg/user"
	"strconv"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}
type UserLoginInfo struct {
	Name string `json: name`
	Pwd  string `json: pwd`
}
type UserModifiedInfo struct {
	Oldpwd string `json: oldpwd`
	Newpwd string `json: newpwd`
}

func login(user_service *account_service.User) int {
	code := e.SUCCESS

	if user_service.Username == "" || user_service.Password == "" {
		code = e.INVALID_PARAMS
		return code
	}

	isRight, err := user_service.Check()
	if err != nil {
		code = e.ERROR
		return code
	}
	if !isRight {
		code = e.INVALID_PARAMS
		return code
	}

	return code
}
func AuthInfo(ctx *context.Context) {
	var code int
	uname, _ := ctx.Get("uname")
	uid, _ := ctx.Get("uid")
	utype, _ := ctx.Get("utype")

	uuid, _ := strconv.Atoi(uid.(string))

	auth_service := auth_service.Auth{Username: uname.(string), UID: uuid}

	_, err := auth_service.Check()

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, "请输入正确的用户名和密码")
		return
	}

	values := map[string]interface{}{
		"uid":     uid,
		"uname":   uname,
		"utype":   utype,
		"auth1":   auth_service.Auth1,
		"auth2":   auth_service.Auth2,
		"auth3":   auth_service.Auth3,
		"succMsg": "获取权限成功",
	}

	ctx.GenResSuccess(values)
	return
}
func AdminLogin(ctx *context.Context) {
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")
	logintype := ctx.PostForm("type")

	user_service := account_service.User{Username: name, Password: pwd}

	var code int

	if logintype == "2" {
		user_service.Usertype = 2
		code = login(&user_service)
	}
	if logintype == "1" {
		user_service.Usertype = 1
		code = login(&user_service)
	}
	if code == e.INVALID_PARAMS {
		ctx.GenResError(code, "请输入正确的用户名和密码")
		return
	}

	if code == e.ERROR {
		ctx.GenResError(code, "")
		return
	}

	if code == e.SUCCESS {
		values := map[string]string{"account": name, "succMsg": "登录成功"}
		token, err := util.GenerateToken(name, logintype, strconv.Itoa(user_service.UID))

		if err != nil {
			code = e.ERROR_AUTH_TOKEN
			ctx.GenResError(code, "")
			return
		} else {
			if setting.ServerSetting.DevMode == "debug" {
				ctx.SetCookie("hualvmall_authorization", token, 30*24*60*60, "/", "127.0.0.1", false, false)
			} else {
				ctx.SetCookie("hualvmall_authorization", token, 30*24*60*60, "/", "hualvmall.com", false, false)
			}
		}

		ctx.GenResSuccess(values)
		return
	}

}

func UserLogin(ctx *context.Context) {

	var userLoginInfo UserLoginInfo
	err := ctx.BindJSON(&userLoginInfo)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "login info 结构不正确")
		return
	}
	user_service := account_service.User{Username: userLoginInfo.Name, Password: userLoginInfo.Pwd, Usertype: 3}
	code := login(&user_service)

	if code == e.INVALID_PARAMS {
		ctx.GenResError(code, "请输入正确的用户名和密码")
		return
	}

	if code == e.ERROR {
		ctx.GenResError(code, "")
		return
	}

	if code == e.SUCCESS {

		token, err := util.GenerateToken(userLoginInfo.Name, "3", strconv.Itoa(user_service.UID))
		expired := 30 * 24 * 60 * 60
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
			ctx.GenResError(code, "")
			return
		} else {
			ctx.SetCookie("hualvmall_staff_authorization", token, expired, "/", "hualvmall.com", false, false)
		}
		expiredStr := strconv.Itoa(expired)

		values := map[string]string{"account": userLoginInfo.Name, "token": token, "expired": expiredStr, "succMsg": "登录成功"}
		ctx.GenResSuccess(values)
		return
	}
}

func AdminCheck(ctx *context.Context) {
	uname, _ := ctx.Get("uname")
	utype, _ := ctx.Get("utype")

	values := map[string]interface{}{"uname": uname, "utype": utype, "succMsg": "登录成功"}
	ctx.GenResSuccess(values)
}
func UserRest(ctx *context.Context) {

	var code int
	uid, _ := strconv.Atoi(ctx.PostForm("uid"))

	var newPwd string
	newPwd = "000000"
	// 加密
	u.CryptoHandler(&newPwd)
	updateAccountValues := map[string]interface{}{"password": newPwd, "salt1": setting.CryptoSetting.Seed1, "salt2": setting.CryptoSetting.Seed2}

	err := dao.UpdateUser(uid, updateAccountValues)

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, err.Error())
		return
	}
	values := map[string]string{"succMsg": "账号密码修改成功"}
	ctx.GenResSuccess(values)

}
func UserRegister(ctx *context.Context) {
	var code int

	auth1 := ctx.Query("fid")
	name := ctx.PostForm("name")
	realname := ctx.PostForm("realname")
	pwd := ctx.PostForm("pwd")

	if name == "" || pwd == "" || realname == "" {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "请输入正确的用户名密码")
		return
	}

	var password = pwd // temp storage

	user_item := account_service.FirmUser{
		Usertype: 4,
		Username: name,
		Password: pwd,
		Realname: realname,
		Auth1:    auth1,
	}
	err := user_item.AccountRegister()

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, err.Error())
		return
	}

	values := map[string]string{"account": name, "password": password, "succMsg": "注册成功"}
	ctx.GenResSuccess(values)
	return
}
func UserList(ctx *context.Context) {
	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))

	total := new(int)
	userlist, err := dao.GetAccountList(total, (pageIndex-1)*pageSize, pageSize, "")

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询用户列表数据失败")
		return
	}

	var userResList []map[string]interface{}

	for _, user_item := range userlist {
		item := make(map[string]interface{})

		item["username"] = user_item.Username
		item["usertype"] = user_item.Usertype
		item["id"] = user_item.ID
		item["realname"] = user_item.Realname

		userResList = append(userResList, item)
	}

	values := map[string]interface{}{"page": "page", "pageSize": "pageSize", "succMsg": "注册成功", "list": userResList, "total": *total}

	ctx.GenResSuccess(values)
}
func UserImport(ctx *context.Context) {
	var code int
	file, info, err := ctx.Request.FormFile("file")

	auth1 := ctx.PostForm("fid")

	if len(auth1) < 1 {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "请选择fid")
		return
	}
	if err != nil {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "请选择文件")
		return
	}

	upload_file_ext := filepath.Ext(info.Filename)

	if upload_file_ext != ".xlsx" {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "请上传正确格式的文件")
		return
	}

	filename := setting.XlsxSetting.UserXlsx

	filename = path.Join("tmp", filename)

	err = tools.Upload(file, filename)
	// 处理文件
	if err == nil {
		err, arr := tools.GenerateDataList(filename)

		if err == nil {

			var accoutList = account_service.ArrayUser{}
			var authList = auth_service.ArrayAuth{}
			var staffList = staff_service.ArrayStaff{}

			for num, row := range arr {
				if num != 0 {
					if err != nil {
						break
					}
					user_item := account_service.User{Usertype: 3, Username: row[0], Password: row[2], Realname: row[1]}
					u.CryptoHandler(&user_item.Password)

					accoutList = append(accoutList, user_item)

				}
			}
			ids, err := accoutList.BuckRegister()

			if err != nil {
				fmt.Println(err)
				code = e.INVALID_PARAMS
				ctx.GenResError(code, err.Error())
				return
			}
			for num, item := range ids {

				auth_item := auth_service.Auth{UID: item, Username: arr[num+1][0], Auth1: auth1, Auth2: "1", Auth3: "1"}

				fid, _ := strconv.Atoi(auth1)
				var staff_item_firm = firm_service.Firm{
					Fid: fid,
				}
				err = staff_item_firm.FindFirm()
				fmt.Println("----------------------")
				fmt.Printf("%+v\n", staff_item_firm)
				if err != nil {
					code = e.INVALID_PARAMS
					ctx.GenResError(code, err.Error())
					return
				}

				staff_item := staff_service.Staff{
					UID:          item,
					Username:     arr[num+1][0],
					Realname:     arr[num+1][1],
					Fid:          staff_item_firm.Fid,
					Firmname:     staff_item_firm.Firmname,
					FirmRealname: staff_item_firm.FirmRealname,
				}
				authList = append(authList, auth_item)
				staffList = append(staffList, staff_item)
			}

			_, err = authList.BuckRegister()

			if err != nil {
				code = e.INVALID_PARAMS
				ctx.GenResError(code, err.Error())
				return
			}

			_, err = staffList.BuckRegister()
			if err != nil {
				code = e.INVALID_PARAMS
				ctx.GenResError(code, err.Error())
				return
			}

			code = e.SUCCESS
			values := map[string]string{"succMsg": "账户批量导入成功, 已存在账户不会被覆盖"}
			ctx.GenResSuccess(values)
			return

		} else {
			code = e.INVALID_PARAMS
			ctx.GenResError(code, "检查xlsx文件格式，默认sheet1放置登录信息")
			return
		}
	}

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, "")
		return
	}

}

func AddUser(ctx *context.Context) {

}

func UserModify(ctx *context.Context) {

	var code int
	uid, _ := ctx.Get("uid")
	UID, _ := strconv.Atoi(uid.(string))

	var userModifiedInfo UserModifiedInfo
	err := ctx.BindJSON(&userModifiedInfo)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "请求包体结构不正确")
		return
	}
	user_service := account_service.User{UID: UID, Password: userModifiedInfo.Oldpwd}

	isRight, err := user_service.CheckPwd()
	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, "密码不正确")
		return
	}
	if !isRight {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "密码不正确")
		return
	}

	if userModifiedInfo.Newpwd == "" {
		code = e.ERROR
		ctx.GenResError(code, "新密码不能为空")
		return
	}
	// 加密
	u.CryptoHandler(&userModifiedInfo.Newpwd)
	updateAccountValues := map[string]interface{}{"password": userModifiedInfo.Newpwd, "salt1": setting.CryptoSetting.Seed1, "salt2": setting.CryptoSetting.Seed2}

	err = dao.UpdateUser(UID, updateAccountValues)

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, err.Error())
		return
	}
	values := map[string]string{"succMsg": "账号密码修改成功"}
	ctx.GenResSuccess(values)

}

func DelegateManager(ctx *context.Context) {
	var code int

	uid, _ := strconv.Atoi(ctx.Query("uid"))
	updateType := ctx.Query("type")

	var updateUserValues map[string]interface{}
	var updateAuthValues map[string]interface{}
	if updateType == "on" {
		updateUserValues = map[string]interface{}{"usertype": 2}
		updateAuthValues = map[string]interface{}{"auth2": "1000", "auth3": "1000"}
	}

	if updateType == "off" {
		updateUserValues = map[string]interface{}{"usertype": 3}
		updateAuthValues = map[string]interface{}{"auth2": "1", "auth3": "1"}
	}
	err := dao.UpdateUser(uid, updateUserValues)

	err = dao.UpdateAuth(uid, updateAuthValues)

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, err.Error())
		return
	}

	values := map[string]string{"succMsg": "管理员设定成功"}
	ctx.GenResSuccess(values)
}

func StaffInfo(ctx *context.Context) {
	url := fmt.Sprintf("/wx/user/info?uid=%v&fid=%v", ctx.Query("uid"), ctx.Query("fid"))
	ctx.Redirect(http.StatusMovedPermanently, url)
}

func StaffOperate(ctx *context.Context) {
	url := fmt.Sprintf("/wx/user/update?uid=%v&fid=%v&data=%v", ctx.Query("uid"), ctx.Query("fid"), ctx.PostForm("data"))
	ctx.Redirect(http.StatusMovedPermanently, url)
}

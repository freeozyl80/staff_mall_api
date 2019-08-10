package user

import (
	"fmt"
	"path"
	"path/filepath"
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/account_service"
	"staff-mall-center/models/service/auth_service"
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

func login(name string, pwd string, logtype string) int {
	code := e.SUCCESS

	if name == "" || pwd == "" {
		code = e.INVALID_PARAMS
		return code
	}

	user_service := account_service.User{Username: name, Password: pwd}

	if logtype == "SuperAdmin" {
		user_service.Usertype = 1
	} else if logtype == "Admin" {
		user_service.Usertype = 2
	} else {
		user_service.Usertype = 3
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

func AdminLogin(ctx *context.Context) {
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")
	logintype := ctx.PostForm("type")
	var code int

	if logintype == "2" {
		code = login(name, pwd, "Admin")
	}
	if logintype == "1" {
		code = login(name, pwd, "SuperAdmin")
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
		token, err := util.GenerateToken(name, logintype)

		if err != nil {
			code = e.ERROR_AUTH_TOKEN
			ctx.GenResError(code, "")
			return
		} else {
			ctx.SetCookie("hualvmall_authorization", token, 30*24*60*60, "/", ".", false, false)
		}

		ctx.GenResSuccess(values)
		return
	}

}

func UserLogin(ctx *context.Context) {
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")

	code := login(name, pwd, "")

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

		token, err := util.GenerateToken(name, "3")

		if err != nil {
			code = e.ERROR_AUTH_TOKEN
			ctx.GenResError(code, "")
			return
		} else {
			ctx.SetCookie("hualvmall_authorization", token, 30*24*60*60, "/", "", false, false)
		}

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

func UserRegister(ctx *context.Context) {
	var code int
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")

	if name == "" || pwd == "" {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "请输入正确的用户名密码")
		return
	}

	var password = pwd // temp storage

	user_service := account_service.User{Usertype: 4, Username: name, Password: pwd}
	isSucc, err := user_service.Register()

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, "")
		return
	}
	if !isSucc {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "注册账户已存在")
		return
	}
	code = e.SUCCESS

	values := map[string]string{"account": name, "password": password, "succMsg": "注册成功"}
	ctx.GenResSuccess(values)
	return
}
func UserList(ctx *context.Context) {
	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))

	userlist, err := dao.GetAccountList((pageIndex-1)*pageSize, pageSize, "")

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

	values := map[string]interface{}{"page": "page", "pageSize": "pageSize", "succMsg": "注册成功", "list": userResList}

	ctx.GenResSuccess(values)
}
func UserImport(ctx *context.Context) {
	var code int

	file, info, err := ctx.Request.FormFile("file")

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
		_err, arr := tools.GenerateUserList(filename)

		if _err == nil {

			var accoutList = account_service.ArrayUser{}
			var authList = auth_service.ArrayAuth{}

			for num, row := range arr {
				if num != 0 {
					usertype, err := strconv.Atoi(row[3])
					if err != nil {
						break
					}
					user_item := account_service.User{Usertype: usertype, Username: row[0], Password: row[2], Realname: row[1]}
					u.CryptoHandler(&user_item.Password)

					accoutList = append(accoutList, user_item)

				}
			}
			ids, uerror := accoutList.BuckRegister()

			if uerror != nil {
				fmt.Println(uerror)
				code = e.INVALID_PARAMS
				ctx.GenResError(code, "检查xlsx文件内容")
				return
			}

			for num, item := range ids {
				usertype, err := strconv.Atoi(arr[num+1][3])
				if err != nil {
					break
				}

				auth_item := auth_service.Auth{UID: item, Username: arr[num+1][0], Auth1: usertype, Auth2: 1}

				authList = append(authList, auth_item)
			}

			_, aeeor := authList.BuckRegister()

			if aeeor != nil {
				code = e.INVALID_PARAMS
				ctx.GenResError(code, "检查xlsx文件内容")
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

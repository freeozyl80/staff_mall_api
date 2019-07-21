package user

import (
	"fmt"
	"path"
	"path/filepath"
	"staff-mall-center/models/service/manage_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"staff-mall-center/pkg/setting"
	"staff-mall-center/pkg/tools"

	"github.com/EDDYCJY/go-gin-example/pkg/util"
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

	user_service := manage_service.User{Username: name, Password: pwd}

	if logtype == "SuperAdmin" {
		user_service.Usertype = 2
	} else if logtype == "Admin" {
		user_service.Usertype = 3
	} else {
		user_service.Usertype = 4
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

	code := login(name, pwd, "Admin")

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
		ctx.GenResSuccess(values)

		return
	}

}

func UserLogin(ctx *context.Context) {
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")

	password := pwd
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

		token, err := util.GenerateToken(name, password)

		if err != nil {
			code = e.ERROR_AUTH_TOKEN
			ctx.GenResError(code, "")
			return
		} else {
			ctx.SetCookie("hualvmall_authorization", token, 30*24*60*60, "/", "", true, false)
		}

		ctx.GenResSuccess(values)
		return
	}
}

func AdminCheck(ctx *context.Context) {

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

	user_service := manage_service.User{Usertype: 4, Username: name, Password: pwd}
	fmt.Println(user_service)
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
			for num, row := range arr {
				if num != 0 {
					user_service := manage_service.User{Usertype: 4, Username: row[0], Password: row[1], Realname: row[2]}
					user_service.Register()
				}
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

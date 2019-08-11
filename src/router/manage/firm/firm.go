package firm

import (
	"staff-mall-center/models/service/firm_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"strconv"
)

func FirmList(ctx *context.Context) {
	values := map[string]interface{}{"page": "page", "pageSize": "pageSize", "succMsg": "注册成功", "list": make([]string, 0)}

	ctx.GenResSuccess(values)
}

func FirmAdd(ctx *context.Context) {
	var code int

	uid, _ := ctx.Get("uid")
	uname, _ := ctx.Get("uname")
	utype, _ := ctx.Get("utype")
	uuid, _ := strconv.Atoi(uid.(string))

	if utype != 2 {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		ctx.GenResError(code, "")
		return
	}

	firmname := ctx.PostForm("firmname")
	firm_realname := ctx.PostForm("firm_realname")
	banlance := ctx.PostForm("banlance")

	if firmname == "" || firm_realname == "" || banlance == "" {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "")
		return
	}

	banlanceInt, err := strconv.Atoi(banlance)

	firm_service := firm_service.Firm{
		UID:          uuid,
		Username:     uname.(string),
		Firmname:     firmname,
		FirmRealname: firm_realname,
		Balance:      banlanceInt,
	}
	isSucc, err := firm_service.CreateFirm()
	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, "")
		return
	}
	if isSucc {
		code = e.SUCCESS

		values := map[string]string{"firmname": firmname, "firm_realname": firm_realname, "succMsg": "注册成功"}
		ctx.GenResSuccess(values)
		return
	}
}

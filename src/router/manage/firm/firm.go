package firm

import (
	"fmt"
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
	uutype, _ := strconv.Atoi(utype.(string))

	if uutype != 1 {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		ctx.GenResError(code, "")
		return
	}

	firmname := ctx.PostForm("firmname")
	firm_realname := ctx.PostForm("firm_realname")
	balance := ctx.PostForm("balance")

	fmt.Println("-----------------", firmname, firm_realname, balance)

	if firmname == "" || firm_realname == "" || balance == "" {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "")
		return
	}
	fmt.Println("-----------------")
	balanceInt, err := strconv.Atoi(balance)

	firm_service := firm_service.Firm{
		UID:          uuid,
		Username:     uname.(string),
		Firmname:     firmname,
		FirmRealname: firm_realname,
		Balance:      balanceInt,
	}
	isSucc, err := firm_service.CreateFirm()
	fmt.Println("-----------------")
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

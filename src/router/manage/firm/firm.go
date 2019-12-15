package firm

import (
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/firm_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"strconv"
)

func FirmList(ctx *context.Context) {

	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))

	firmList, err := dao.GetFirmList((pageIndex-1)*pageSize, pageSize, "")

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询用户列表数据失败")
		return
	}

	var firmResList []map[string]interface{}

	for _, firm_item := range firmList {
		item := make(map[string]interface{})
		item["fid"] = firm_item.ID
		item["firmname"] = firm_item.Firmname
		item["firm_realname"] = firm_item.FirmRealname
		item["balance"] = firm_item.Balance

		firmResList = append(firmResList, item)
	}

	values := map[string]interface{}{"page": pageIndex, "pageSize": pageSize, "succMsg": "查询成功", "list": firmResList}

	ctx.GenResSuccess(values)
}
func StaffList(ctx *context.Context) {
	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	fid, _ := strconv.Atoi(ctx.Query("fid"))

	seachValue := map[string]interface{}{"fid": fid}

	total := new(int)
	userlist, err := dao.GetStaffList(total, (pageIndex-1)*pageSize, pageSize, seachValue)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询用户列表数据失败")
		return
	}

	values := map[string]interface{}{"page": "page", "pageSize": "pageSize", "succMsg": "查询成功", "list": userlist, "total": *total}

	ctx.GenResSuccess(values)
}
func AccountList(ctx *context.Context) {

	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	auth1, _ := strconv.Atoi(ctx.Query("fid"))

	seachValue := map[string]interface{}{"auth1": auth1}

	total := new(int)
	userlist, err := dao.GetAuthList(total, (pageIndex-1)*pageSize, pageSize, seachValue)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询用户列表数据失败")
		return
	}

	var userResList []map[string]interface{}

	for _, user_item := range userlist {
		item := make(map[string]interface{})

		item["username"] = user_item.Username
		item["id"] = user_item.ID
		item["uid"] = user_item.UID
		item["auth2"] = user_item.Auth2

		userResList = append(userResList, item)
	}

	values := map[string]interface{}{"page": "page", "pageSize": "pageSize", "succMsg": "查询成功", "list": userResList, "total": *total}

	ctx.GenResSuccess(values)
}
func FirmAdd(ctx *context.Context) {
	var code int

	utype, _ := ctx.Get("utype")
	uutype, _ := strconv.Atoi(utype.(string))

	if uutype != 1 {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		ctx.GenResError(code, "")
		return
	}

	firmname := ctx.PostForm("firmname")
	firm_realname := ctx.PostForm("firm_realname")
	balance := ctx.PostForm("balance")

	if firmname == "" || firm_realname == "" || balance == "" {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "")
		return
	}
	balanceInt, err := strconv.Atoi(balance)

	firm_service := firm_service.Firm{
		Firmname:     firmname,
		FirmRealname: firm_realname,
		Balance:      balanceInt,
	}
	isSucc, err := firm_service.CreateFirm()

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, err.Error())
		return
	}
	if isSucc {
		code = e.SUCCESS

		values := map[string]string{"firmname": firmname, "firm_realname": firm_realname, "succMsg": "注册成功"}
		ctx.GenResSuccess(values)
		return
	}
}
func FirmDetail(ctx *context.Context) {

	Fid, _ := strconv.Atoi(ctx.Query("fid"))

	var firm_item = firm_service.Firm{
		Fid: Fid,
	}
	err := firm_item.FindFirm()

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询公司数据失败")
		return
	}
	firmResInfo := make(map[string]interface{})

	firmResInfo["firm_name"] = firm_item.Firmname
	firmResInfo["product_group"] = firm_item.ProductGroup
	firmResInfo["category_group"] = firm_item.CategoryGroup
	firmResInfo["firm_realname"] = firm_item.FirmRealname
	firmResInfo["balance"] = firm_item.Balance

	values := map[string]interface{}{"succMsg": "查询成功", "info": firmResInfo}

	ctx.GenResSuccess(values)
}
func FirmUpdate(ctx *context.Context) {

	Fid, _ := strconv.Atoi(ctx.Query("fid"))

	var firm_item = firm_service.Firm{
		Fid: Fid,
	}
	err := firm_item.FindFirm()

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询公司数据失败")
		return
	}

	var code int

	utype, _ := ctx.Get("utype")
	uutype, _ := strconv.Atoi(utype.(string))

	if uutype != 1 {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		ctx.GenResError(code, "")
		return
	}

	firmname := ctx.PostForm("firmname")
	firm_realname := ctx.PostForm("firm_realname")
	balance := ctx.PostForm("balance")
	product_group := ctx.PostForm("product_group")
	category_group := ctx.PostForm("category_group")

	balanceInt, err := strconv.Atoi(balance)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "balance 是 使用 数字 类型")
	}
	var updateFirmValues map[string]interface{}
	if len(product_group) > 0 {
		updateFirmValues = map[string]interface{}{
			"balance":        balanceInt,
			"firm_name":      firmname,
			"firm_realname":  firm_realname,
			"product_group":  product_group,
			"category_group": category_group,
		}
	} else {
		updateFirmValues = map[string]interface{}{
			"balance": balanceInt, "firm_name": firmname, "firm_realname": firm_realname}
	}

	err = dao.UpdateFirm(Fid, updateFirmValues)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "更新失败")
	}

	values := map[string]string{"firmname": firmname, "firm_realname": firm_realname, "succMsg": "注册成功"}

	ctx.GenResSuccess(values)
}

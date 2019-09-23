package staff

import (
	"encoding/json"
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/staff_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"strconv"
)

type StaffUpdateStruct struct {
	Coin int `json:coin`
}

func UserInfo(ctx *context.Context) {
	uid, isExist := ctx.Get("uid")

	var UID int
	var FID int

	if !isExist {
		UID, _ = strconv.Atoi(ctx.Query("uid"))
		FID, _ = strconv.Atoi(ctx.Query("fid"))

	} else {
		UID, _ = strconv.Atoi(uid.(string))
	}

	var staff_item = staff_service.Staff{
		UID: UID,
	}
	err := staff_item.GetInfo()

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询用户信息失败"+err.Error())
		return
	}

	staffResItem := map[string]interface{}{
		"sid":          staff_item.SID,
		"uid":          staff_item.UID,
		"username":     staff_item.Username,
		"realname":     staff_item.Realname,
		"fid":          staff_item.Fid,
		"firmname":     staff_item.Firmname,
		"gender":       staff_item.Gender,
		"user_address": staff_item.UserAddress,
		"user_avatar":  staff_item.UserAvatar,
		"tel":          staff_item.Tel,
		"birthday":     staff_item.Birthday,
		"coi":          staff_item.Coin,
	}
	if !isExist {
		if FID != staff_item.Fid {
			code := e.ERROR_AUTH_CHECK_TOKEN_FAIL
			ctx.GenResError(code, "所属员工与公司不符，无法查看")
			return
		}
	}
	values := map[string]interface{}{"succMsg": "查询成功", "info": staffResItem}

	ctx.GenResSuccess(values)
}

func UpdateStaffInfo(ctx *context.Context) {
	var UID int

	UID, _ = strconv.Atoi(ctx.Query("uid"))
	DATA := ctx.Query("data")

	var updateStaffInfo StaffUpdateStruct
	err := json.Unmarshal([]byte(DATA), &updateStaffInfo)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "update info 结构不正确")
		return
	}

	updateStaffValues := map[string]interface{}{"Coin": updateStaffInfo.Coin}

	err = dao.UpdateStaffItem(UID, updateStaffValues)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}

	values := map[string]interface{}{"succMsg": "staff 更新成功"}

	ctx.GenResSuccess(values)
}

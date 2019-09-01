package staff

import (
	"staff-mall-center/models/service/staff_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"strconv"
)

func UserInfo(ctx *context.Context) {
	uid := ctx.Param("uid")

	UID, _ := strconv.Atoi(uid)

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
		"fid":          staff_item.FID,
		"firmname":     staff_item.Firmname,
		"gender":       staff_item.Gender,
		"user_address": staff_item.UserAddress,
		"user_avatar":  staff_item.UserAvatar,
		"tel":          staff_item.Tel,
		"birthday":     staff_item.Birthday,
		"coi ":         staff_item.Coin,
	}

	values := map[string]interface{}{"succMsg": "查询成功", "info": staffResItem}

	ctx.GenResSuccess(values)
}

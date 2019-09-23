package staff

import (
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/firm_service"
	"staff-mall-center/models/service/staff_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"strconv"
	"strings"
)

func UpdateStaffCoin(ctx *context.Context) {
	var UID int
	var FID int
	var Coin int

	UID, _ = strconv.Atoi(ctx.Query("uid"))
	FID, _ = strconv.Atoi(ctx.Query("fid"))
	Coin, _ = strconv.Atoi(ctx.PostForm("data"))

	var staff_item = staff_service.Staff{
		UID: UID,
	}
	err := staff_item.GetInfo()

	var firm_item = firm_service.Firm{
		Fid: FID,
	}
	err = firm_item.FindFirm()

	// 判断是否够费用：
	if firm_item.Balance < Coin {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "公司剩余福利点数不够")
		return
	}
	err = dao.UpdateStaffCoin(UID, FID, Coin, staff_item.Coin, firm_item.Balance)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}

	values := map[string]interface{}{"succMsg": "staff 更新成功"}

	ctx.GenResSuccess(values)
}
func UpdateStaffListCoin(ctx *context.Context) {
	var FID int
	var Coin int

	UIDS := strings.Split(ctx.Query("uids"), ",")
	FID, _ = strconv.Atoi(ctx.Query("fid"))
	Coin, _ = strconv.Atoi(ctx.PostForm("data"))

	var firm_item = firm_service.Firm{
		Fid: FID,
	}
	err := firm_item.FindFirm()
	// 判断是否够费用：
	if firm_item.Balance < Coin*len(UIDS) {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "公司剩余福利点数不够")
		return
	}

	for _, val := range UIDS {
		UID, _ := strconv.Atoi(val)

		var staff_item = staff_service.Staff{
			UID: UID,
		}
		var firm_item = firm_service.Firm{
			Fid: FID,
		}

		err = firm_item.FindFirm()
		err = staff_item.GetInfo()
		err = dao.UpdateStaffCoin(UID, FID, Coin, staff_item.Coin, firm_item.Balance)
	}

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}

	values := map[string]interface{}{"succMsg": "福利点数 更新成功"}

	ctx.GenResSuccess(values)
}

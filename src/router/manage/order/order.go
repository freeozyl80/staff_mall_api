package order

import (
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/order_service"
	"staff-mall-center/models/service/staff_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

func OrderList(ctx *context.Context) {

	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	fid, _ := strconv.Atoi(ctx.Query("fid"))

	seachValue := map[string]interface{}{"Fid": fid}

	orderList, err := dao.GetOrderList((pageIndex-1)*pageSize, pageSize, seachValue)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询商品列表数据失败")
		return
	}

	var orderResList []map[string]interface{}

	for _, order_item := range orderList {
		item := make(map[string]interface{})

		item["order_id"] = order_item.OrderID
		item["order_status"] = order_item.OrderStatus
		item["order_product_info"] = order_item.ProductInfo
		item["order_total_price"] = order_item.ProductTotalPrice
		item["order_receiving_username"] = order_item.ReceivingUsername
		item["order_receiving_tel"] = order_item.ReceivingUserTel
		item["order_receiving_city"] = order_item.ReceivingUserCity
		item["order_receiving_address"] = order_item.ReceivingUserAddress

		orderResList = append(orderResList, item)
	}

	values := map[string]interface{}{"page": "page", "pageSize": "pageSize", "succMsg": "查询成功", "list": orderResList}

	ctx.GenResSuccess(values)
}

func CancelOrder(ctx *context.Context) {
	fid, _ := strconv.Atoi(ctx.Query("fid"))
	orderid, err := uuid.FromString(ctx.PostForm("orderId"))

	var order_item = order_service.Order{
		OrderID: orderid,
	}
	err = order_item.GetOrderInfo()

	var staff_item = staff_service.Staff{
		UID: order_item.UID,
	}
	err = staff_item.GetInfo()

	err = dao.Refund(fid, orderid, order_item.ProductTotalPrice, order_item.UID, staff_item.Coin)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询商品列表数据失败")
		return
	}

	values := map[string]interface{}{"succMsg": "订单取消 更新成功"}

	ctx.GenResSuccess(values)
}

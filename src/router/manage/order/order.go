package order

import (
	"fmt"
	"net/http"
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

	var seachValue map[string]interface{}
	if fid != 0 {
		seachValue = map[string]interface{}{"Fid": fid}
	} else {
		seachValue = map[string]interface{}{}
	}

	total := new(int)
	orderList, err := dao.GetOrderList(total, (pageIndex-1)*pageSize, pageSize, seachValue)

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
		item["fid"] = order_item.Fid
		item["uid"] = order_item.UID

		orderResList = append(orderResList, item)
	}

	values := map[string]interface{}{"page": "page", "pageSize": "pageSize", "succMsg": "查询成功", "list": orderResList, "total": *total}

	ctx.GenResSuccess(values)
}
func OrderDetail(ctx *context.Context) {
	url := fmt.Sprintf("/wx/order/item/%v?uid=%v", ctx.Query("order_id"), ctx.Query("order_id"))
	ctx.Redirect(http.StatusMovedPermanently, url)
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

	err = dao.Refund(fid, orderid, order_item.ProductTotalPrice, order_item.UID, order_item.ProductInfo, staff_item.Coin)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询商品列表数据失败")
		return
	}

	values := map[string]interface{}{"succMsg": "订单取消 更新成功"}

	ctx.GenResSuccess(values)
}
func DeliverOrder(ctx *context.Context) {

	orderid, err := uuid.FromString(ctx.PostForm("orderId"))

	var order_item = order_service.Order{
		OrderID: orderid,
	}

	err = order_item.GetOrderInfo()

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询订单列表数据失败")
		return
	}

	if order_item.OrderStatus != 2 {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}
	updateOrderValues := map[string]interface{}{"OrderStatus": 3}

	err = dao.UpdateOrderItem(order_item.UID, order_item.OrderID, updateOrderValues)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}

	values := map[string]interface{}{"succMsg": "订单发货 更新成功"}

	ctx.GenResSuccess(values)
}

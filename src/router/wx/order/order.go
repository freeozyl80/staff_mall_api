package order

import (
	"encoding/json"
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/order_service"
	"staff-mall-center/models/service/product_service"
	"staff-mall-center/models/service/staff_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"strconv"
)

type OrderInfo struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}
type ReceivingInfo struct {
	Username    string `json:"username"`
	Tel         int    `json:"tel"`
	UserAddress string `json:"user_address"`
	UserCity    string `json:"user_city"`
}
type OrderRes struct {
	ProductList   []OrderInfo   `json:"product_list"`
	ReceivingInfo ReceivingInfo `json:"receiving_info"`
}

func GenerateOrder(ctx *context.Context) {
	var reqInfo OrderRes
	err := ctx.BindJSON(&reqInfo)

	prouctList := reqInfo.ProductList
	addressInfo := reqInfo.ReceivingInfo

	var productResList []OrderInfo
	var productTotalPrice int
	for _, item := range prouctList {
		var product_item = product_service.Product{
			PID: item.Id,
		}
		err = product_item.Occupy(item.Count)
		if err == nil {
			var orderItem = OrderInfo{
				Id:    product_item.PID,
				Count: item.Count,
			}

			productResList = append(productResList, orderItem)
			productTotalPrice += product_item.ProductPrice * item.Count
		}
	}
	productInfo, err := json.Marshal(productResList)

	uid, _ := ctx.Get("uid")
	UID, _ := strconv.Atoi(uid.(string))

	var staff_item = staff_service.Staff{
		UID: UID,
	}
	err = staff_item.GetInfo()

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询用户信息失败"+err.Error())
		return
	}
	order_service := order_service.Order{
		OrderStatus: 1,

		UID:      staff_item.UID,
		Username: staff_item.Username,
		Realname: staff_item.Realname,
		Tel:      staff_item.Tel,

		ReceivingUsername:    addressInfo.Username,
		ReceivingUserTel:     addressInfo.Tel,
		ReceivingUserCity:    addressInfo.UserCity,
		ReceivingUserAddress: addressInfo.UserAddress,

		ProductTotalPrice: productTotalPrice,
		ProductInfo:       string(productInfo),
	}

	err = order_service.Generate()

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "订单失败"+err.Error())
		return
	}

	values := map[string]interface{}{"succMsg": "提交订单成功", "orderId": order_service.OrderID}

	ctx.GenResSuccess(values)
}

func ListOrder(ctx *context.Context) {
	uid, _ := ctx.Get("uid")
	UID, _ := strconv.Atoi(uid.(string))

	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))

	seachValue := map[string]interface{}{"UID": UID}

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

package order

import (
	"encoding/json"
	"fmt"
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/order_service"
	"staff-mall-center/models/service/product_service"
	"staff-mall-center/models/service/staff_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type OrderInfo struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}

type OrderInfoList []OrderInfo

type UnusualOrderInfo struct {
	Id          int `json:"id"`
	Count       int `json:"count"`
	OriginCount int `json:origin_count`
}

type ReceivingInfo struct {
	Username    string `json:"username"`
	Tel         string `json:"tel"`
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
	var unusualProductResList []UnusualOrderInfo
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
		} else if err.Error() == "OverRemain" {
			// 超过最大数量，count 为 0
			var unusualOrderItem = UnusualOrderInfo{
				Id:          product_item.PID,
				Count:       product_item.ProductCount,
				OriginCount: item.Count,
			}

			fmt.Printf("%+v\n", unusualOrderItem)
			unusualProductResList = append(unusualProductResList, unusualOrderItem)
		}
	}
	// 所有订单库存都不够，所有直接返回，不生产订单
	if len(productResList) == 0 {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "生成订单失败，请核数据")
		return
	}

	productInfo, err := json.Marshal(productResList)
	var unusualProductInfo string
	if len(unusualProductResList) != 0 {
		unusualProductInfoByte, _ := json.Marshal(unusualProductResList)
		unusualProductInfo = string(unusualProductInfoByte)
	}

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

		Fid: staff_item.Fid,

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

	values := map[string]interface{}{"succMsg": "提交订单成功", "order_id": order_service.OrderID, "order_info": string(productInfo), "unusual_product_info": unusualProductInfo}

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

		var order_product_info_list OrderInfoList
		err := json.Unmarshal([]byte(order_item.ProductInfo), &order_product_info_list)

		if err != nil {
			code := e.INVALID_PARAMS
			ctx.GenResError(code, "查询商品列表数据失败(解析产品详情错误)")
			return
		}

		var product_item []map[string]interface{}
		for _, product := range order_product_info_list {

			prod_item := product_service.Product{
				PID: product.Id,
			}
			err := prod_item.Find()

			if err != nil {
				code := e.INVALID_PARAMS
				ctx.GenResError(code, err.Error())
				return
			}

			pitem := make(map[string]interface{})
			pitem["product_id"] = product.Id
			pitem["count"] = product.Count
			pitem["product_name"] = prod_item.ProductName
			pitem["product_realname"] = prod_item.ProductRealname
			pitem["category_id"] = prod_item.CategoryID
			pitem["category_name"] = prod_item.CategoryName
			pitem["category_realname"] = prod_item.CategoryRealname
			pitem["supplier_name"] = prod_item.SupplierName
			pitem["supplier_realname"] = prod_item.SupplierRealname
			pitem["product_price"] = prod_item.ProductPrice
			pitem["product_count"] = prod_item.ProductCount
			pitem["product_img"] = prod_item.ProductImg
			pitem["product_status"] = prod_item.ProductStatus
			pitem["product_desc"] = prod_item.ProductDesc
			product_item = append(product_item, pitem)
		}

		item["order_id"] = order_item.OrderID
		item["order_status"] = order_item.OrderStatus
		item["order_product_info"] = product_item //order_item.ProductInfo
		item["order_total_price"] = order_item.ProductTotalPrice
		item["order_receiving_username"] = order_item.ReceivingUsername
		item["order_receiving_tel"] = order_item.ReceivingUserTel
		item["order_receiving_city"] = order_item.ReceivingUserCity
		item["order_receiving_address"] = order_item.ReceivingUserAddress
		item["order_created_time"] = order_item.CreatedOn

		orderResList = append(orderResList, item)
	}

	values := map[string]interface{}{"page": "page", "pageSize": "pageSize", "succMsg": "查询成功", "list": orderResList}

	ctx.GenResSuccess(values)
}

func GetOrderInfo(ctx *context.Context) {
	uid, _ := ctx.Get("uid")
	UID, _ := strconv.Atoi(uid.(string))

	orderID := ctx.Param("id")

	OrderID, err := uuid.FromString(orderID)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "order Id 不存在")
		return
	}

	var order_item = order_service.Order{
		UID:     UID,
		OrderID: OrderID,
	}
	err = order_item.GetOrderInfo()

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}

	var order_product_info_list OrderInfoList
	err = json.Unmarshal([]byte(order_item.ProductInfo), &order_product_info_list)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询商品列表数据失败(解析产品详情错误)")
		return
	}

	var product_item []map[string]interface{}
	for _, product := range order_product_info_list {

		prod_item := product_service.Product{
			PID: product.Id,
		}
		err := prod_item.Find()

		if err != nil {
			code := e.INVALID_PARAMS
			ctx.GenResError(code, err.Error())
			return
		}

		pitem := make(map[string]interface{})
		pitem["product_id"] = product.Id
		pitem["count"] = product.Count
		pitem["product_name"] = prod_item.ProductName
		pitem["product_realname"] = prod_item.ProductRealname
		pitem["category_id"] = prod_item.CategoryID
		pitem["category_name"] = prod_item.CategoryName
		pitem["category_realname"] = prod_item.CategoryRealname
		pitem["supplier_name"] = prod_item.SupplierName
		pitem["supplier_realname"] = prod_item.SupplierRealname
		pitem["product_price"] = prod_item.ProductPrice
		pitem["product_count"] = prod_item.ProductCount
		pitem["product_img"] = prod_item.ProductImg
		pitem["product_status"] = prod_item.ProductStatus
		pitem["product_desc"] = prod_item.ProductDesc
		product_item = append(product_item, pitem)
	}
	oitem := make(map[string]interface{})
	oitem["order_id"] = order_item.OrderID
	oitem["order_status"] = order_item.OrderStatus
	oitem["order_product_info"] = product_item //order_item.ProductInfo
	oitem["order_total_price"] = order_item.ProductTotalPrice
	oitem["order_receiving_username"] = order_item.ReceivingUsername
	oitem["order_receiving_tel"] = order_item.ReceivingUserTel
	oitem["order_receiving_city"] = order_item.ReceivingUserCity
	oitem["order_receiving_address"] = order_item.ReceivingUserAddress
	oitem["order_created_time"] = order_item.CreatedTime

	values := map[string]interface{}{"succMsg": "订单查询成功", "info": oitem}

	ctx.GenResSuccess(values)
}

func OperateOrder(ctx *context.Context) {
	uid, _ := ctx.Get("uid")
	UID, _ := strconv.Atoi(uid.(string))
	orderID := ctx.Param("id")
	OrderID, err := uuid.FromString(orderID)

	orderAimStatus := ctx.Param("status")
	OrderAimStatus, _ := strconv.Atoi(orderAimStatus)

	var msg string
	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "order Id 不存在")
		return
	}
	var order_item = order_service.Order{
		UID:     UID,
		OrderID: OrderID,
	}
	fmt.Printf("%+v\n", order_item)
	err = order_item.GetOrderInfo()

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}
	if OrderAimStatus == 5 {
		msg = "订单已取消"
		code := e.INVALID_PARAMS
		if !(order_item.OrderStatus == 1 || order_item.OrderStatus == 2) {
			ctx.GenResError(code, "该订单无法取消")
			return
		}
	}

	if OrderAimStatus == 2 {
		msg = "订单已支付"
		code := e.INVALID_PARAMS
		if order_item.OrderStatus != 1 {
			ctx.GenResError(code, "该订单无法支付")
			return
		}
		var staff_item = staff_service.Staff{
			UID: UID,
		}
		err := staff_item.GetInfo()

		if err != nil {
			ctx.GenResError(code, "用户查询失败")
			return
		}
		if staff_item.Coin-order_item.ProductTotalPrice < 0 {
			ctx.GenResError(code, "用户福利点数不足支付")
			return
		}

		updateStaffValues := map[string]interface{}{"Coin": staff_item.Coin - order_item.ProductTotalPrice}
		err = dao.UpdateStaffItem(UID, updateStaffValues)
		if err != nil {
			ctx.GenResError(code, "福利点数扣去失败")
			return
		}
	}

	updateOrderValues := map[string]interface{}{"OrderStatus": OrderAimStatus}

	err = dao.UpdateOrderItem(UID, OrderID, updateOrderValues)

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}

	values := map[string]interface{}{"succMsg": msg, "order_id": order_item.OrderID}

	ctx.GenResSuccess(values)
}

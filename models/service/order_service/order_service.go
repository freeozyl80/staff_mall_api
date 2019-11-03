package order_service

import (
	"fmt"
	"staff-mall-center/models/dao"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	ID int

	OrderID     uuid.UUID
	OrderStatus int

	Fid int

	ProductInfo       string
	ProductTotalPrice int

	UID                  int
	Username             string
	Realname             string
	Tel                  string
	ReceivingUsername    string
	ReceivingUserTel     string
	ReceivingUserAddress string
	ReceivingUserCity    string
}

func (o *Order) Generate() error {

	order, err := dao.GenerateOrder(
		o.OrderStatus,
		o.ProductInfo,
		o.ProductTotalPrice,
		o.UID,
		o.Username,
		o.Realname,
		o.Tel,
		o.Fid,
		o.ReceivingUsername,
		o.ReceivingUserCity,
		o.ReceivingUserTel,
		o.ReceivingUserAddress,
	)

	o.OrderID = order.OrderID

	return err
}

func (o *Order) GetOrderInfo() error {
	_order, err := dao.GetOrderItem(o.UID, o.OrderID)
	fmt.Printf("%+v\n", _order)
	if err != nil {
		return err
	}

	o.OrderID = _order.OrderID

	o.OrderStatus = _order.OrderStatus
	o.ProductInfo = _order.ProductInfo
	o.ProductInfo = _order.ProductInfo
	o.ProductTotalPrice = _order.ProductTotalPrice

	o.Username = _order.Username
	o.Realname = _order.Realname
	o.Tel = _order.Tel

	o.ReceivingUsername = _order.ReceivingUsername
	o.ReceivingUserTel = _order.ReceivingUserTel
	o.ReceivingUserAddress = _order.ReceivingUserAddress
	o.ReceivingUserCity = _order.ReceivingUserCity

	return err
}

package order_service

import (
	"staff-mall-center/models/dao"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	ID int

	OrderID     uuid.UUID
	OrderStatus int

	ProductInfo       string
	ProductTotalPrice int

	UID                  int
	Username             string
	Realname             string
	Tel                  int
	ReceivingUsername    string
	ReceivingUserTel     int
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
		o.ReceivingUsername,
		o.ReceivingUserCity,
		o.ReceivingUserTel,
		o.ReceivingUserAddress,
	)

	o.OrderID = order.OrderID

	return err
}

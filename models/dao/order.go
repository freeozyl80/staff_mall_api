package dao

import (
	"errors"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Order struct {
	Model
	OrderID     uuid.UUID `gorm:"not null;"`
	OrderStatus int       `gorm:"not null;" json:"order_status" comment '1: 下定， 2.代表支付成，3.发货完成，4. 收货完成， 5. 取消，6，退换中， 7异常'`

	ProductInfo       string `gorm:"not null;" json:"product_info"`
	ProductTotalPrice int    `gorm:"not null;" json:"product_total_price"`

	UID      int    `gorm:"not null;" json:"uid"`
	Username string `gorm:"not null;" json:"username"`
	Realname string `gorm:"not null;" json:"realname"`
	Tel      int    `gorm:"not null;" json:"tel"`

	ReceivingUsername    string `gorm:"not null;" json:"receiving_username"`
	ReceivingUserTel     int    `gorm:"not null;" json:"receiving_tel"`
	ReceivingUserCity    string `gorm:"not null;" json:"receiving_user_city"`
	ReceivingUserAddress string `gorm:"not null;" json:"receiving_user_address"`
}

func (base *Order) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("OrderID", uuid)
}

func GenerateOrder(
	order_status int,
	product_info string,
	product_total_price int,

	uid int,
	username string,
	realname string,
	tel int,
	receiving_username string,
	receiving_user_city string,
	receiving_tel int,
	receiving_user_address string,
) (Order, error) {

	var order = Order{
		OrderStatus: order_status,

		ProductInfo:       product_info,
		ProductTotalPrice: product_total_price,
		UID:               uid,
		Username:          username,
		Realname:          realname,
		Tel:               tel,

		ReceivingUsername:    receiving_username,
		ReceivingUserTel:     receiving_tel,
		ReceivingUserCity:    receiving_user_city,
		ReceivingUserAddress: receiving_user_address,
	}

	if err := db.Create(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}
func GetOrderList(pageIndex int, pageSize int, maps interface{}) ([]*Order, error) {
	var orderList []*Order
	err := db.Where(maps).Offset(pageIndex).Limit(pageSize).Find(&orderList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return orderList, nil
}

func GetOrderItem(uid int, orderId uuid.UUID) (Order, error) {
	var order Order

	err := db.Where(Order{UID: uid, OrderID: orderId}).First(&order).Error

	if order.ID > 0 && err == nil {
		return order, nil
	} else {
		return order, errors.New("can not find")
	}
}

func UpdateOrderItem(uid int, orderId uuid.UUID, data interface{}) error {

	if err := db.Model(&Order{}).Where(Order{UID: uid, OrderID: orderId}).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Order struct {
	Model
	OrderID     uuid.UUID `gorm:"not null;"`
	OrderStatus int       `gorm:"not null;" json:"order_status"`

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
	fmt.Printf("%+v\n", maps)
	err := db.Where(maps).Offset(pageIndex).Limit(pageSize).Find(&orderList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return orderList, nil
}

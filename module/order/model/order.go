package model

import (
	"errors"
	"strings"

	modelProduct "pizza-order/module/product/model"
	modelUser "pizza-order/module/user/model"

	"gorm.io/gorm"
)

var (
	ErrAddressCannotBeEmpty  = errors.New("address cannot be empty")
)

const (
	EntityName = "Order"
)

type OrderStatusType string

var OrderStatus = struct {
    SUCCESS         OrderStatusType
    PENDING OrderStatusType
    FAILED         OrderStatusType
    DELIVERY OrderStatusType
}{
    SUCCESS:         "SUCCESS",
    PENDING: "PENDING",
    FAILED:         "FAILED",
    DELIVERY: "DELIVERY",
}

type Order struct {
	gorm.Model
	UserID  int
	User modelUser.User
	Address string `gorm:"size:200" form:"address"`
	Status string `gorm:"size:20; default:PENDING"`
	OrderProducts []OrderProduct
}

type OrderProduct struct {
	OrderID  uint `gorm:"primaryKey"`
	ProductID int `gorm:"primaryKey"`
	Quantity int
	Price int

	Product modelProduct.Product
}

type OrderCreation struct {
	UserID  int `form:"user_id"`
	Address string `form:"address"`
	Products []OrderProduct
}

type OrderProductCreation struct {
	ProductID int
	Quantity int
	Price int
}

func (i *OrderCreation) Validate() error {
	i.Address = strings.TrimSpace(i.Address)

	if i.Address == "" {
		return ErrAddressCannotBeEmpty
	}

	return nil
}

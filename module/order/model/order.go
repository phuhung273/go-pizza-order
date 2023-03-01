package model

import (
	"errors"
	"strings"

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
	UserID  int32
	Address string `gorm:"size:200" form:"address"`
	Status string `gorm:"size:20; default:PENDING"`
	// TODO: user, products relation
}

type OrderCreation struct {
	UserID  string `form:"user_id"`
	Address string `form:"address"`
}

func (i *OrderCreation) Validate() error {
	i.Address = strings.TrimSpace(i.Address)

	if i.Address == "" {
		return ErrAddressCannotBeEmpty
	}

	return nil
}

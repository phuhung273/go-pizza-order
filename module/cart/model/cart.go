package model

import (
	"errors"
)

var (
	ErrProductIdCannotBeEmpty  = errors.New("product Id cannot be empty")
)

const (
	EntityName = "Order"
)

type Cart struct {
	Items []*CartItem
}

type CartItem struct {
	ID int32 `json:"id"`
	Quantity int32
}

func (i *CartItem) Validate() error {
	if i.ID == 0 {
		return ErrProductIdCannotBeEmpty
	}

	return nil
}

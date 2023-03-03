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
	Quantity int
	Price int
}

type CartItem struct {
	ID int `json:"id"`
	Name string
	Image string
	Quantity int
	Price int
}

func (i *CartItem) Validate() error {
	if i.ID == 0 {
		return ErrProductIdCannotBeEmpty
	}

	return nil
}

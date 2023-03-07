package model

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

var (
	ErrNameCannotBeEmpty  = errors.New("name cannot be empty")
	ErrImageCannotBeEmpty = errors.New("image cannot be empty")
	ErrPriceCannotBeEmpty = errors.New("price cannot be empty")
)

const (
	EntityName = "Product"
)

type Product struct {
	gorm.Model
	Name  string `gorm:"size:50" form:"name"`
	Image string `gorm:"size:200" form:"image"`
	Price int  `form:"price"`
}

type ProductCreation struct {
	Name  string `form:"name"`
	Image string `form:"image"`
	Price int  `form:"price"`
}

func (i *ProductCreation) Validate() error {
	i.Name = strings.TrimSpace(i.Name)
	i.Image = strings.TrimSpace(i.Image)

	if i.Name == "" {
		return ErrNameCannotBeEmpty
	} else if i.Image == "" {
		return ErrImageCannotBeEmpty
	} else if i.Price == 0 {
		return ErrPriceCannotBeEmpty
	}

	return nil
}

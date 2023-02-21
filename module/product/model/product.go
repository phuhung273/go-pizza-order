package model

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

var (
	ErrNameCannotBeEmpty  = errors.New("Name cannot be empty")
	ErrImageCannotBeEmpty = errors.New("Image cannot be empty")
	ErrPriceCannotBeEmpty = errors.New("Price cannot be empty")
)

type Product struct {
	gorm.Model
	Name  string `gorm:"size:50"`
	Image string `gorm:"size:200"`
	Price int32
}

func (i *Product) Validate() error {
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

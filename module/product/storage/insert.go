package storage

import (
	"context"
	"pizza-order/module/product/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.ProductCreation) error {
	dbData := model.Product{
		Name:  data.Name,
		Image: data.Image,
		Price: data.Price,
	}
	if err := s.db.Create(&dbData).Error; err != nil {
		return err
	}

	return nil
}

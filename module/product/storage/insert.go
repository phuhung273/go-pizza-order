package storage

import (
	"context"
	"pizza-order/module/product/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.Product) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

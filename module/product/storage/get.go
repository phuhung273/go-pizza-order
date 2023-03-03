package storage

import (
	"context"
	"pizza-order/module/product/model"
)

func (s *sqlStore) GetItem(ctx context.Context, id int) (model.Product, error) {
	dbData := model.Product{}
	if err := s.db.First(&dbData, id).Error; err != nil {
		return model.Product{}, err
	}

	return dbData, nil
}

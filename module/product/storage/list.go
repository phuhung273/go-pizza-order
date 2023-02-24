package storage

import (
	"context"
	"pizza-order/common"
	"pizza-order/module/product/model"
)

func (s *sqlStore) ListItem(
	ctx context.Context,
) ([]model.Product, error) {
	var result []model.Product

	if err := s.db.Find(&result).Error; err != nil {

		return nil, common.ErrDB(err)
	}

	return result, nil
}

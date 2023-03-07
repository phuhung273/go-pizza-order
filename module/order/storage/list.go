package storage

import (
	"context"
	"pizza-order/common"
	"pizza-order/module/order/model"
)

func (s *sqlStore) ListItem(
	ctx context.Context,
) ([]model.Order, error) {
	var result []model.Order

	if err := s.db.Preload("OrderProducts.Product").Preload("User").Find(&result).Error; err != nil {

		return nil, common.ErrDB(err)
	}

	return result, nil
}

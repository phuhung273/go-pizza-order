package storage

import (
	"context"
	"pizza-order/module/order/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.OrderCreation) error {
	dbData := model.Order{
		UserID: data.UserID,
		Address:  data.Address,
		Status: string(model.OrderStatus.PENDING),
		OrderProducts: data.Products,
	}

	if err := s.db.Create(&dbData).Error; err != nil {
		return err
	}

	return nil
}

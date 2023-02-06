package storage

import (
	"context"
	"pizza-order/module/user/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.User) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

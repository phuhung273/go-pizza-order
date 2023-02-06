package storage

import (
	"context"
	"pizza-order/module/user/model"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.User, error) {
	var data model.User

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

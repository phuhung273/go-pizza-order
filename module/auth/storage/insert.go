package storage

import (
	"context"
	authModel "pizza-order/module/auth/model"
	userModel "pizza-order/module/user/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *authModel.RegisterRequest) error {
	dbData := userModel.User{}
	dbData.Username = data.Username
	dbData.Password = data.Password

	if err := s.db.Create(&dbData).Error; err != nil {
		return err
	}

	return nil
}

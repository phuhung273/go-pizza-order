package biz

import (
	"context"
	"pizza-order/module/auth/model"

	"golang.org/x/crypto/bcrypt"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.RegisterRequest) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.RegisterRequest) error {
	if err := data.Validate(); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	data.Password = string(hash)

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}

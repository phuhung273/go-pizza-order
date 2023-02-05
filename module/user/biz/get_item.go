package biz

import (
	"context"
	"pizza-order/module/user/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.User, error)
}

type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItemByUsername(ctx context.Context, username string) (*model.User, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

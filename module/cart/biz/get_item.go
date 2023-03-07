package biz

import (
	"context"
	"pizza-order/common"
	"pizza-order/module/cart/model"
)

type GetItemStorage interface {
	GetItem(
		ctx context.Context,
	) (model.Cart, error)
}

type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItem(
	ctx context.Context,
) (model.Cart, error) {

	data, err := biz.store.GetItem(ctx)

	if err != nil {
		return model.Cart{}, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}

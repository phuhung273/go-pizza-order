package biz

import (
	"context"
	"pizza-order/common"
	"pizza-order/module/product/model"
)

type GetItemStorage interface {
	GetItem(
		ctx context.Context,
		id int,
	) (model.Product, error)
}

type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItem(
	ctx context.Context,
	id int,
) (model.Product, error) {

	data, err := biz.store.GetItem(ctx, id)

	if err != nil {
		return model.Product{}, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}

package biz

import (
	"context"
	"pizza-order/common"
	"pizza-order/module/order/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
	) ([]model.Order, error)
}

type listItemBiz struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *listItemBiz {
	return &listItemBiz{store: store}
}

func (biz *listItemBiz) ListItem(
	ctx context.Context,
) ([]model.Order, error) {

	data, err := biz.store.ListItem(ctx)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}

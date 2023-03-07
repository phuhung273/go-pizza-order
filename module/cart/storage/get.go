package storage

import (
	"context"
	"pizza-order/common"
	"pizza-order/module/cart/model"
)

func (s *sessionStore) GetItem(ctx context.Context) (model.Cart, error) {
	store := *s.db;
	item, ok := store.Get("cart").(model.Cart)
	
	if !ok {
		return model.Cart{}, common.RecordNotFound
	}
	

	return item, nil
}
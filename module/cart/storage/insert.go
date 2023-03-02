package storage

import (
	"context"
	"pizza-order/module/cart/model"
)

func (s *sessionStore) CreateItem(ctx context.Context, data *model.CartItem) error {
	store := *s.db;
	cart, ok := store.Get("cart").(model.Cart)
	
	if !ok {
		cart = model.Cart{
			Items: []*model.CartItem{
				{
					ID: data.ID,
					Quantity: 1,
				},
			},
		}
	} else {
		for i, v := range cart.Items {
			if v.ID == data.ID {
				cart.Items[i].Quantity++
				break
			}
		}
	}

	store.Set("cart", cart)
	err := store.Save()
	if err != nil {
		return err
	}

	return nil
}

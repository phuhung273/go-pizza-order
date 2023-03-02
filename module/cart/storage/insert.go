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
			Quantity: 0, // Will plus 1 below
		}
	} else {
		existed := false;

		for i, v := range cart.Items {
			if v.ID == data.ID {
				cart.Items[i].Quantity = cart.Items[i].Quantity + 1
				existed = true
				break
			}
		}

		if !existed {
			cart.Items = append(cart.Items, &model.CartItem{
				ID: data.ID,
				Quantity: 1,
			})
		}
	}

	cart.Quantity = cart.Quantity + 1;

	store.Set("cart", cart)
	if err := store.Save(); err != nil {
		return err
	}

	return nil
}

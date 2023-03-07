package ginorder

import (
	"net/http"
	modelAuth "pizza-order/module/auth/model"
	bizCart "pizza-order/module/cart/biz"
	storageCart "pizza-order/module/cart/storage"
	"pizza-order/module/order/biz"
	"pizza-order/module/order/model"
	"pizza-order/module/order/storage"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var requestData model.OrderCreation

		if err := c.ShouldBind(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		if err := requestData.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		session := sessions.Default(c)
		storeCart := storageCart.NewSessionStore(&session)

		getBusiness := bizCart.GetItemStorage(storeCart)

		cart, err := getBusiness.GetItem(c.Request.Context());

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		auth, ok := session.Get("auth").(modelAuth.AuthSession)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No user id",
			})

			return
		}

		requestData.UserID = auth.ID
		for _, v := range cart.Items {
			requestData.Products = append(requestData.Products, model.OrderProduct{
				ProductID: v.ID,
				Quantity: v.Quantity,
				Price: v.Price,
			})
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)

		if err := business.CreateNewItem(c.Request.Context(), &requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "done",
		})
	}
}

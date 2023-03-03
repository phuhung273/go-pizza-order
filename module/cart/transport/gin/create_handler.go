package ginproduct

import (
	"net/http"
	"pizza-order/module/cart/biz"
	"pizza-order/module/cart/model"
	"pizza-order/module/cart/storage"

	bizProduct "pizza-order/module/product/biz"
	storageProduct "pizza-order/module/product/storage"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var requestData model.CartItem

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
		store := storage.NewSessionStore(&session)
		business := biz.NewCreateItemBiz(store)

		productStore := storageProduct.NewSQLStore(db)
		productBiz := bizProduct.NewGetItemBiz(productStore);

		product, err := productBiz.GetItem(c.Request.Context(), requestData.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		requestData.Price = product.Price
		requestData.Name = product.Name
		requestData.Image = product.Image

		if err := business.CreateNewItem(c.Request.Context(), &requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		getBusiness := biz.GetItemStorage(store)

		data, err := getBusiness.GetItem(c.Request.Context());

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"totalQty": data.Quantity,
		})
	}
}

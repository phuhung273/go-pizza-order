package ginproduct

import (
	"net/http"
	"pizza-order/module/cart/biz"
	"pizza-order/module/cart/model"
	"pizza-order/module/cart/storage"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateItem() func(ctx *gin.Context) {
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

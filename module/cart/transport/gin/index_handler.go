package ginproduct

import (
	"net/http"
	"pizza-order/module/cart/biz"
	"pizza-order/module/cart/storage"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		store := storage.NewSessionStore(&session)

		getBusiness := biz.GetItemStorage(store)

		data, err := getBusiness.GetItem(c.Request.Context());

		if err != nil {
			c.HTML(http.StatusBadRequest, "cart.html", gin.H{
				"error": err.Error(),
			})

			return
		}

		c.HTML(http.StatusOK, "cart.html", gin.H{
			"data": data,
		})
	}
}

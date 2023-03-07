package ginorder

import (
	"net/http"
	"pizza-order/module/order/biz"
	"pizza-order/module/order/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		store := storage.NewSQLStore(db)
		business := biz.NewListItemBiz(store)

		result, err := business.ListItem(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func Index(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "orders.html", gin.H{})
	}
}

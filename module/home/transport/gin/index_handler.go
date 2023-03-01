package ginhome

import (
	"net/http"
	"pizza-order/module/product/biz"
	"pizza-order/module/product/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(db *gorm.DB) func(ctx *gin.Context) {
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

		c.HTML(http.StatusOK, "home.html", gin.H{
			"data": result,
		})
	}
}

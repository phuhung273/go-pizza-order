package ginproduct

import (
	"net/http"
	"pizza-order/module/product/biz"
	"pizza-order/module/product/model"
	"pizza-order/module/product/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var requestData model.Product

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

		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)

		if err := business.CreateNewItem(c.Request.Context(), &requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.String(http.StatusOK, "register done")
	}
}

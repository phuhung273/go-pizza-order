package ginauth

import (
	"net/http"
	authBiz "pizza-order/module/auth/biz"
	authModel "pizza-order/module/auth/model"
	userBiz "pizza-order/module/user/biz"
	userModel "pizza-order/module/user/model"
	"pizza-order/module/user/storage"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var requestData authModel.AuthRequest

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

		itemData := userModel.User{
			Username: requestData.Username,
			Password: requestData.Password,
		}

		store := storage.NewSQLStore(db)
		business := authBiz.NewCreateItemBiz(store)

		if err := business.CreateNewItem(c.Request.Context(), &itemData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.String(http.StatusOK, "register done")
	}
}

func Login(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var requestData authModel.AuthRequest

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
		business := userBiz.NewGetItemBiz(store)

		user, err := business.GetItemByUsername(c.Request.Context(), requestData.Username)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestData.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.String(http.StatusOK, "login done")
	}
}

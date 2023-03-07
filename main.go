package main

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	modelAuth "pizza-order/module/auth/model"
	ginauth "pizza-order/module/auth/transport/gin"
	modelCart "pizza-order/module/cart/model"
	gincart "pizza-order/module/cart/transport/gin"
	ginhome "pizza-order/module/home/transport/gin"
	modelOrder "pizza-order/module/order/model"
	ginorder "pizza-order/module/order/transport/gin"
	modelProduct "pizza-order/module/product/model"
	ginproduct "pizza-order/module/product/transport/gin"
	modelUser "pizza-order/module/user/model"
)

func main() {
	godotenv.Load()

	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()
	log.Println("DB Connection:", db)

	// Auto Migrate
	db.AutoMigrate(&modelUser.User{}, &modelProduct.Product{}, &modelOrder.Order{}, &modelOrder.OrderProduct{})
	
	r := gin.Default()

	gob.Register(modelCart.Cart{})
	gob.Register(modelAuth.AuthSession{})

	store := cookie.NewStore([]byte("secret"))
  	r.Use(sessions.Sessions("mysession", store))

	r.Static("/public", "./public")
	r.Static("/favicon.ico", "./public/favicon.ico")

	r.SetFuncMap(template.FuncMap{
        "mul": func (a int, b int) int {
			return a * b;
		},
    })
	r.LoadHTMLGlob("views/**/*")

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "register.html", gin.H{})
	})

	r.GET("/", ginhome.Index(db))

	v1 := r.Group("/v1")
	{
		items := v1.Group("/auth")
		{
			items.POST("/register", ginauth.Register(db))
			items.POST("/login", ginauth.Login(db))
		}

		products := v1.Group("/product")
		{
			products.POST("/", ginproduct.CreateItem(db))
		}

		carts := v1.Group("/cart")
		{
			carts.POST("/", gincart.CreateItem(db))
		}

		orders := v1.Group("/order")
		{
			orders.GET("/", ginorder.List(db))
			orders.POST("/", ginorder.CreateItem(db))
		}

	}
	
	r.GET("/orders", ginorder.Index(db))
	r.GET("/cart", gincart.Index())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT"))); err != nil {
		log.Fatalln(err)
	}
}

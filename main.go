package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	ginauth "pizza-order/module/auth/transport/gin"
	ginproduct "pizza-order/module/product/transport/gin"
	"pizza-order/module/user/model"
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
	db.AutoMigrate(&model.User{})

	r := gin.Default()

	r.Static("/public", "./public")
	r.Static("/favicon.ico", "./public/favicon.ico")
	r.LoadHTMLGlob("views/**/*")

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "register.html", gin.H{})
	})

	v1 := r.Group("/v1")
	{
		items := v1.Group("/auth")
		{
			items.POST("/register", ginauth.Register(db))
			items.POST("/login", ginauth.Login(db))
		}

		products := v1.Group("/products")
		{
			products.POST("/", ginproduct.CreateItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT"))); err != nil {
		log.Fatalln(err)
	}
}

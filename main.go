package main

import (
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"pizza-order/module/auth/transport/gin"
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

	v1 := r.Group("/v1")
	{
		items := v1.Group("/auth")
		{
			items.POST("/register", ginauth.Register(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	ginPost "social-photo/modules/post/transport/gin"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		posts := v1.Group("/posts")
		{
			posts.GET("/")
			posts.GET("/:id", ginPost.GetPostById(db))
			posts.POST("/", ginPost.CreatePost(db))
			posts.PATCH("/:id")
			posts.DELETE("/:id")
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"social-photo/component/tokenprovider/jwt"
	"social-photo/middleware"
	ginPost "social-photo/modules/post/transport/gin"
	"social-photo/modules/user/storage"
	ginUser "social-photo/modules/user/transport/gin"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	secretKey := os.Getenv("SECRET_KEY")

	if err != nil {
		log.Fatal(err)
	}

	authStore := storage.NewSQLStore(db)
	tokenProvider := jwt.NewTokenProvider("jwt", secretKey)
	middlewareAuth := middleware.RequiredAuth(authStore, tokenProvider)

	r := gin.Default()
	r.Use(middleware.Recovery())

	v1 := r.Group("/v1")
	{
		v1.POST("/register", ginUser.Register(db))
		v1.POST("/login", ginUser.Login(db, tokenProvider))
		v1.GET("/profile", middlewareAuth, ginUser.Profile())

		posts := v1.Group("/posts", middlewareAuth)
		{
			posts.GET("/", ginPost.ListPost(db))
			posts.GET("/:id", ginPost.GetPostById(db))
			posts.POST("/", ginPost.CreatePost(db))
			posts.PATCH("/:id", ginPost.UpdatePost(db))
			posts.DELETE("/:id", ginPost.DeletePostById(db))
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

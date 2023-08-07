package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"social-photo/common"
	"social-photo/component/tokenprovider/jwt"
	"social-photo/component/uploadprovider"
	"social-photo/middleware"
	ginPost "social-photo/modules/post/transport/gin"
	ginUpload "social-photo/modules/upload/transport/gin"
	"social-photo/modules/user/storage"
	ginUser "social-photo/modules/user/transport/gin"
	ginLikePost "social-photo/modules/userlikepost/transport/gin"
	"social-photo/pubsub"
	"social-photo/subscriber"
)

func main() {
	config, err := common.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:")
	}

	dsn := os.Getenv(config.DB_CONN_STR)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	secretKey := os.Getenv(config.JWT_SECRET_KEY)

	// S3_Provider
	s3BucketName := os.Getenv(config.S3BucketName)
	s3Region := os.Getenv(config.S3Region)
	s3APIKey := os.Getenv(config.S3APIKey)
	s3SecretKey := os.Getenv(config.S3APISecret)
	s3Domain := os.Getenv(config.S3Domain)

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	if err != nil {
		log.Fatal(err)
	}

	authStore := storage.NewSQLStore(db)
	tokenProvider := jwt.NewTokenProvider("jwt", secretKey)
	middlewareAuth := middleware.RequiredAuth(authStore, tokenProvider)

	r := gin.Default()
	r.Use(middleware.Recovery())

	// pub/sub
	ps := pubsub.NewPubSub()
	_ = subscriber.NewEngine(db, ps).Start()

	v1 := r.Group("/v1")
	{

		// upload
		v1.POST("/upload", ginUpload.Upload(s3Provider))

		// user
		v1.POST("/register", ginUser.Register(db))
		v1.POST("/login", ginUser.Login(db, tokenProvider))
		v1.GET("/profile", middlewareAuth, ginUser.Profile())

		// post
		posts := v1.Group("/posts")
		{
			posts.GET("/", ginPost.ListPost(db))
			posts.GET("/:id", ginPost.GetPostById(db))
			posts.POST("/", middlewareAuth, ginPost.CreatePost(db))
			posts.PATCH("/:id", middlewareAuth, ginPost.UpdatePost(db))
			posts.DELETE("/:id", middlewareAuth, ginPost.DeletePostById(db))

			// like post
			posts.POST("/:id/like", middlewareAuth, ginLikePost.LikePost(db, ps))
			posts.POST("/:id/unlike", middlewareAuth, ginLikePost.UnlikePost(db, ps))
			posts.GET("/:id/like", middlewareAuth, ginLikePost.ListUserLiked(db))
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

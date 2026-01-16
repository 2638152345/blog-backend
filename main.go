package main

import (
	"github.com/gin-gonic/gin"

	"blog-backend/config"
	"blog-backend/handlers"
	"blog-backend/middleware"
	"blog-backend/models"
)

func main() {
	config.InitDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)

	config.InitLogger()
	config.Logger.Println("Service started")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.GET("/posts", handlers.GetPosts)
	r.GET("/posts/:id", handlers.GetPost)
	r.GET("/posts/:post_id/comments", handlers.GetComments)

	auth := r.Group("/")
	auth.Use(middleware.JWTAUTHMIDDLEWARE())
	{
		auth.POST("/posts", handlers.CreatePost)
		auth.PUT("/posts/:id", handlers.UpdatePost)
		auth.DELETE("/posts/:id", handlers.DeletePost)

		auth.POST("/comments", handlers.CreateComment)
	}
	r.Run(":8080")
}

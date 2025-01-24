package server

import (
	v1 "example.com/m/server/v1"
	"example.com/m/storage"
	"github.com/gin-gonic/gin"
)

type Options struct {
	Strg storage.StorageI
}

func NewServer(opts *Options) *gin.Engine {
	router := gin.New()

	//ROUTER
	handler := v1.New(&v1.HandlerV1{
		Strg: opts.Strg,
	})

	//Users
	router.POST("/v1/users", handler.CreateUser)
	router.DELETE("/v1/user/:id", handler.DeleteUser)
	router.GET("/v1/user/:id", handler.GetUser)
	router.PUT("/v1/user/:id", handler.UpdateUser)

	//Posts
	router.POST("/v1/posts", handler.CreatePost)
	router.PUT("/v1/post/:id", handler.UpdatePost)
	router.GET("/v1/post/:id", handler.GetPost)
	router.DELETE("/v1/post/:id", handler.DeletePost)

	return router
}

package main

import (
	"github.com/gin-gonic/gin"
	"go-categories-microservice/cmd/api/controllers"
	"go-categories-microservice/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	categorysRoutes := router.Group("/categories")
	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	categorysRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

	categorysRoutes.GET("/", func(ctx *gin.Context) {
		controllers.ListCategory(ctx, inMemoryCategoryRepository)
	})
}

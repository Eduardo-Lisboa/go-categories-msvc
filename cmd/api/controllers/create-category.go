package controllers

import (
	"github.com/gin-gonic/gin"
	"go-categories-microservice/internal/repositories"
	use_cases "go-categories-microservice/internal/use-cases"
	"net/http"
)

type createCategory struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategory

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"succes": false,
				"erro":   err,
			})
		return
	}

	useCase := use_cases.NewCreateCategoryUseCase(repository)
	err := useCase.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"succes": false,
				"error":  err.Error(),
			})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"succes": true})
}

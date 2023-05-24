package controllers

import (
	"github.com/gin-gonic/gin"
	"go-categories-microservice/internal/repositories"
	use_cases "go-categories-microservice/internal/use-cases"
	"net/http"
)

func ListCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {

	useCase := use_cases.NewListeCategoryUseCase(repository)

	categories, err := useCase.Execute()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"succes": false,
				"error":  err.Error(),
			})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"succes": true, "categories": categories})
}

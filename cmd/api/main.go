package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	CategoryRoutes(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

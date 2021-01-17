package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create gin router
	router := gin.Default()

	SetRoutes(router)

	// Launch Gin and
	// handle potential error
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

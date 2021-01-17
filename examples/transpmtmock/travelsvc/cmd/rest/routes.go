package main

import "github.com/gin-gonic/gin"

func SetRoutes(r gin.IRouter) {
	r.GET("/tripsvcflow/:tap", tripSvcflowGetH)
	r.POST("/tripsvcflow", tripSvcflowPostH)
}

package main

import "github.com/gin-gonic/gin"

func SetRoutes(r gin.IRouter) {
	r.GET("/tripsvcflow/", tripSvcflowGetH)
	r.POST("/tripsvcflow", tripSvcflowPostH)
}

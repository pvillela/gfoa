package main

import (
	"github.com/gin-gonic/gin"
)

// func TripSvcPostHandler0(c *gin.Context) {
// 	// Create empty request body
// 	// struct used to bind actual body into
// 	tripRequest := &rpc.TripRequest{}

// 	// Bind JSON content of request body to
// 	// struct created above
// 	err := c.BindJSON(tripRequest)
// 	if err != nil {
// 		// Gin automatically returns an error
// 		// response when the BindJSON operation
// 		// fails, we simply have to stop this
// 		// function from continuing to execute

// 		fmt.Println(err)

// 		c.JSON(400, gin.H{
// 			"msg":  err.Error(),
// 			"code": 999,
// 		})
// 		return
// 	}

// 	// Create response object
// 	tripResponse := tripSvc(*tripRequest)
// 	body := &tripResponse

// 	// And send it off to the requesting client
// 	c.JSON(http.StatusOK, body)
// }

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

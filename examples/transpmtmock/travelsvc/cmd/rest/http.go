package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/internal/boot"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
)

var config = boot.TravelConfig{
	CfgForValidateTripRequestBf:  "CfgForValidateTripRequestBf",
	CfgForGetPostedCardStateDaf:  "CfgForGetPostedCardStateDaf",
	CfgForGetUnpostedUsagesDaf:   "CfgForGetUnpostedUsagesDaf",
	CfgForUpdateCardStateBf:      "CfgForUpdateCardStateBf",
	CfgForRateTripSc:             "CfgForRateTripSc",
	CfgForPrepareUsageBf:         "CfgForPrepareUsageBf",
	CfgForWriteUsageDaf:          "CfgForWriteUsageDaf",
	CfgForPrepareDownstreamEvtBf: "CfgForPrepareDownstreamEvtBf",
	CfgForDownstreamEp:           "CfgForDownstreamEp",
	CfgForPrepareResponseBf:      "CfgForPrepareResponseBf",
}

var tripSvc = boot.TripSvcflowboot(config)

func TripSvcGetHandler(c *gin.Context) {
	// Create response object
	body := &rpc.TripResponse{
		Resp: c.Param("tap"),
	}

	// Send it off to the client
	c.JSON(http.StatusOK, body)
}

func TripSvcPseudoPostHandler(fillReqCb func(interface{}) error) (interface{}, error) {
	// Create empty request body
	// struct used to bind actual body into
	pReq := &rpc.TripRequest{}

	err := fillReqCb(pReq)
	if err != nil {
		return nil, err
	}

	resp := tripSvc(*pReq)

	return &resp, nil
}

type PseudoPostHandler = func(func(interface{}) error) (interface{}, error)

func PostHanderMaker(pseudoHdlr PseudoPostHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		filler := func(pReq interface{}) error {
			// Bind JSON content of request body to
			// struct created above
			err := c.BindJSON(pReq)
			if err != nil {
				// Gin automatically returns an error
				// response when the BindJSON operation
				// fails, we simply have to stop this
				// function from continuing to execute

				fmt.Println(err)

				c.JSON(400, gin.H{
					"msg":  err.Error(),
					"code": 999,
				})
				return err
			}
			return nil
		}
		pResp, err := pseudoHdlr(filler)

		if err != nil {
			return
		}

		c.JSON(http.StatusOK, pResp)
	}
}

func TripSvcPostHandler0(c *gin.Context) {
	// Create empty request body
	// struct used to bind actual body into
	tripRequest := &rpc.TripRequest{}

	// Bind JSON content of request body to
	// struct created above
	err := c.BindJSON(tripRequest)
	if err != nil {
		// Gin automatically returns an error
		// response when the BindJSON operation
		// fails, we simply have to stop this
		// function from continuing to execute

		fmt.Println(err)

		c.JSON(400, gin.H{
			"msg":  err.Error(),
			"code": 999,
		})
		return
	}

	// Create response object
	tripResponse := tripSvc(*tripRequest)
	body := &tripResponse

	// And send it off to the requesting client
	c.JSON(http.StatusOK, body)
}

func main() {
	// Create gin router
	router := gin.Default()

	// Set up GET endpoint
	// for route  /users/<username>
	router.GET("/tripsvcflow/:tap", TripSvcGetHandler)

	// Set up POST endpoint for route /users/<usernane>
	// router.POST("/tripsvcflow", TripSvcPostHandler0)
	TripSvcPostHandler := PostHanderMaker(TripSvcPseudoPostHandler)
	router.POST("/tripsvcflow", TripSvcPostHandler)

	// Launch Gin and
	// handle potential error
	err := router.Run(":8003")
	if err != nil {
		panic(err)
	}
}

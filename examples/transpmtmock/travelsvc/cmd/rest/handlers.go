package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
	"github.com/pvillela/gfoa/pkg/web/wgin"
)

func tripSvcflowGetH(c *gin.Context) {
	// Create response object
	body := &rpc.TripResponse{
		Resp: c.Param("tap"),
	}

	// Send it off to the client
	c.JSON(http.StatusOK, body)
}

var tripSvcflowPostH = wgin.PostHanderMaker(&rpc.TripRequest{}, tripSvcflowP)

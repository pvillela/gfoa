package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
	"github.com/pvillela/gfoa/pkg/web"
	"github.com/pvillela/gfoa/pkg/web/wgin"
)

func tripSvcGetHandler(c *gin.Context) {
	// Create response object
	body := &rpc.TripResponse{
		Resp: c.Param("tap"),
	}

	// Send it off to the client
	c.JSON(http.StatusOK, body)
}

func TripSvcPseudoPostHandler(fillReqCb web.Filler) (interface{}, error) {
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

var tripSvcPostHandler = wgin.PostHanderMaker(&rpc.TripRequest{}, tripSvcAny)

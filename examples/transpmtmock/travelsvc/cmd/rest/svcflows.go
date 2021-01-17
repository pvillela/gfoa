package main

import (
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/internal/boot"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
)

var tripSvc = boot.TripSvcflowboot(config)

func tripSvcAny(pReq interface{}) interface{} {
	tripSvc := boot.TripSvcflowboot(config)
	return tripSvc(*pReq.(*rpc.TripRequest))
}

package main

import (
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/internal/boot"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
)

func tripSvcflowP(pInput interface{}) interface{} {
	tripSvc := boot.TripSvcflowboot(config)
	return tripSvc(*pInput.(*rpc.TripRequest))
}

package main

import (
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
	"github.com/pvillela/gfoa/pkg/web/wgin"
)

var tripSvcflowGetH = wgin.GeneralGetHanderMaker(tripSvcflowM)

var tripSvcflowPostH = wgin.PostHanderMaker(&rpc.TripRequest{}, tripSvcflowP)

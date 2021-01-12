package main

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/internal/boot"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
)

func main() {
	config := boot.TravelConfig{
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

	tripRequest := rpc.TripRequest{
		CardInfo:   "card123",
		DeviceInfo: "device999",
	}

	tripSvc := boot.TripSvcflowboot(config)

	tripResponse := tripSvc(tripRequest)

	fmt.Println("#####", tripResponse)
}

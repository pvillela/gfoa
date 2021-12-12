/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

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

	tripResponse, _ := tripSvc(tripRequest)

	fmt.Println("#####", tripResponse)
}

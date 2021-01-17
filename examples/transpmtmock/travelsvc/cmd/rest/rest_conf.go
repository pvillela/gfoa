package main

import "github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/internal/boot"

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

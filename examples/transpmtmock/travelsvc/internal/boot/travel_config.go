/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package boot

// TravelConfig contains configuration properties for the service.
type TravelConfig struct {
	CfgForValidateTripRequestBf  string
	CfgForGetPostedCardStateDaf  string
	CfgForGetUnpostedUsagesDaf   string
	CfgForUpdateCardStateBf      string
	CfgForRateTripSc             string
	CfgForPrepareUsageBf         string
	CfgForWriteUsageDaf          string
	CfgForPrepareDownstreamEvtBf string
	CfgForDownstreamEp           string
	CfgForPrepareResponseBf      string
}

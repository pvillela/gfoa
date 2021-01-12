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

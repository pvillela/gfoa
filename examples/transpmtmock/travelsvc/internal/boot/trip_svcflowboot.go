package boot

import "github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/internal/fn"

// TripSvcflowboot produces an fn.TripSvcflow from configuration
func TripSvcflowboot(config TravelConfig) fn.TripSvcflow {
	s := fn.TripSvcflowS{
		ValidateTripRequestBf: fn.ValidateTripRequestBfS{
			Cfg: config.CfgForValidateTripRequestBf,
		}.Make(),
		GetUpdatedCardStateFlow: GetUpdatedCardStateFlowboot(config),
		RateTripSc: fn.RateTripScS{
			Cfg: config.CfgForRateTripSc,
		}.Make(),
		PrepareUsageBf: fn.PrepareUsageBfS{
			Cfg: config.CfgForPrepareUsageBf,
		}.Make(),
		WriteUsageDaf: fn.WriteUsageDafS{
			Cfg: config.CfgForWriteUsageDaf,
		}.Make(),
		PrepareDownstreamEvtBf: fn.PrepareDownstreamEvtBfS{
			Cfg: config.CfgForPrepareDownstreamEvtBf,
		}.Make(),
		DownstreamEp: fn.DownstreamEpS{
			Cfg: config.CfgForDownstreamEp,
		}.Make(),
		PrepareResponseBf: fn.PrepareResponseBfS{
			Cfg: config.CfgForPrepareResponseBf,
		}.Make(),
	}
	return s.Make()
}

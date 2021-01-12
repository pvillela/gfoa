package boot

import "github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/internal/fn"

// GetUpdatedCardStateFlowboot produces an fn.GetUpdatedCardStateFlow from configuration
func GetUpdatedCardStateFlowboot(config TravelConfig) fn.GetUpdatedCardStateFlow {
	s := fn.GetUpdatedCardStateFlowS{
		GetPostedCardStateDaf: fn.GetPostedCardStateDafS{
			Cfg: config.CfgForGetPostedCardStateDaf,
		}.Make(),
		GetUnpostedUsagesDaf: fn.GetUnpostedUsagesDafS{
			Cfg: config.CfgForGetUnpostedUsagesDaf,
		}.Make(),
		UpdateCardStateBf: fn.UpdateCardStateBfS{
			Cfg: config.CfgForUpdateCardStateBf,
		}.Make(),
	}
	return s.Make()
}

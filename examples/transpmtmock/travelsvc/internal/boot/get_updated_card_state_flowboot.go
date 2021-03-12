/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

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

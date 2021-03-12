/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package fn

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"
)

// PrepareUsageBfS contains the dependencies required for the construction of a
// PrepareUsageBf.
type PrepareUsageBfS struct {
	Cfg string
}

// PrepareUsageBf is the type of a funtion that creates a Usage from a CardInfo,
// a DeviceInfo, a CardState, and a RateTripResult.
type PrepareUsageBf = func(model.CardInfo, model.DeviceInfo, model.CardState, RateTripResult) model.Usage

// Make creates a PrepareUsageBf from its dependencies.
func (s PrepareUsageBfS) Make() PrepareUsageBf {
	var funcTypeExemplar PrepareUsageBfS
	return func(cardInfo model.CardInfo, deviceInfo model.DeviceInfo, cardState model.CardState, rateResult RateTripResult) model.Usage {
		str := StrApply(funcTypeExemplar, s.Cfg, string(cardInfo), string(deviceInfo), string(cardState), string(rateResult))
		fmt.Println("***", str)
		return model.Usage(str)
	}
}

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package fn

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
)

// PrepareResponseBfS contains the dependencies required for the construction of a
// PrepareResponseBf.
type PrepareResponseBfS struct {
	Cfg string
}

// PrepareResponseBf is the type of a funtion that creates a TripResponse from a CardState
// and a RateTripResult.
type PrepareResponseBf = func(model.CardState, RateTripResult) rpc.TripResponse

// Make creates a PrepareResponseBf from its dependencies.
func (s PrepareResponseBfS) Make() PrepareResponseBf {
	var funcTypeExemplar PrepareResponseBfS
	return func(cardState model.CardState, rateTripResult RateTripResult) rpc.TripResponse {
		str := StrApply(funcTypeExemplar, s.Cfg, string(cardState), string(rateTripResult))
		fmt.Println("***", str)
		return rpc.TripResponse{Resp: str}
	}
}

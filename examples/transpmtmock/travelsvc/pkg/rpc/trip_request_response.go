/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package rpc

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"
)

// TripRequest is the request data structure for the trip service.
type TripRequest struct {
	CardInfo   model.CardInfo   `json:"cardInfo"`
	DeviceInfo model.DeviceInfo `json:"deviceInfo"`
}

// ToString returns a string represenation.
func (req TripRequest) ToString() string {
	return fmt.Sprintf("%v", req)
}

// TripResponse is the response of the trip service.
type TripResponse struct {
	Resp string `json:"resp"`
}

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package main

import (
	"errors"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/internal/boot"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
	log "github.com/sirupsen/logrus"
)

var tripSvc = boot.TripSvcflowboot(config)

func tripSvcflowM(m map[string]string) (interface{}, error) {
	cardInfo, cardInfoOk := m["cardInfo"]
	deviceInfo, deviceInfoOk := m["deviceInfo"]

	log.Info("m[cardInfo]", m["cardInfo"])
	log.Info("m[deviceInfo]", m["deviceInfo"])

	errMsg := ""
	if !cardInfoOk {
		errMsg = "cardInfo parameter not found"
	}
	if !deviceInfoOk {
		if errMsg != "" {
			errMsg = errMsg + ", "
		}
		errMsg = errMsg + "deviceInfo parameter not found"
	}

	var err error
	if errMsg != "" {
		err = errors.New(errMsg)
		return rpc.TripResponse{}, err
	}

	input := rpc.TripRequest{
		CardInfo:   model.CardInfo(cardInfo),
		DeviceInfo: model.DeviceInfo(deviceInfo),
	}

	return tripSvc(input)
}

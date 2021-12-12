/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package main

import (
	"github.com/pvillela/gfoa/pkg/web"
	"github.com/pvillela/gfoa/pkg/web/wgin"
)

var tripSvcflowGetH = wgin.SimpleMapGetHandlerMaker(tripSvcflowM)

var tripSvcflowPostH = wgin.PostHandlerMaker(tripSvc, web.DefaultErrorHandler)

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package fn

import "github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"

// WriteUsageDafS contains the dependencies required for the construction of a
// WriteUsageDaf.
type WriteUsageDafS struct {
	Cfg string
}

// WriteUsageDaf is the type of a funtion that writes a Usage to the database.
type WriteUsageDaf = func(model.Usage)

// Make creates a WriteUsageDaf from its dependencies.
func (s WriteUsageDafS) Make() func(model.Usage) {
	return func(model.Usage) {
	}
}

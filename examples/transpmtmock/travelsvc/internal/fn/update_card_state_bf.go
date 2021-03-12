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

// UpdateCardStateBfS contains the dependencies required for the construction of a
// UpdateCardStateBf.
type UpdateCardStateBfS struct {
	Cfg string
}

// UpdateCardStateBf is the type of a funtion that creates an updated card state in
// memory based on the latest posted card state and unposted usages (previously retrieved from
// the database).
type UpdateCardStateBf = func(model.CardState, []model.Usage) model.CardState

// Make creates a UpdateCardStateBf from its dependencies.
func (s UpdateCardStateBfS) Make() UpdateCardStateBf {
	var funcTypeExemplar UpdateCardStateBfS
	return func(cardStateFromDb model.CardState, usages []model.Usage) model.CardState {
		var strUsages []string
		for _, usage := range usages {
			strUsages = append(strUsages, string(usage))
		}
		args := append([]string{string(cardStateFromDb)}, strUsages...)
		str := StrApply(funcTypeExemplar, s.Cfg, args...)
		fmt.Println("***", str)
		return model.CardState(str)
	}
}

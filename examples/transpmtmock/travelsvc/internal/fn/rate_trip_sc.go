package fn

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"
)

// RateTripScS contains the dependencies required for the construction of a
// RateTripSc.
type RateTripScS struct {
	Cfg string
}

// RateTripSc is the type of a funtion that calls the trip rating service.
type RateTripSc = func(model.CardState, model.CardInfo, model.DeviceInfo) RateTripResult

// RateTripResult is the result type of RateTripSc.
type RateTripResult string

// Make creates a RateTripSc from its dependencies.
func (s RateTripScS) Make() RateTripSc {
	var funcTypeExemplar RateTripScS
	return func(cardState model.CardState, cardInfo model.CardInfo, devInfo model.DeviceInfo) RateTripResult {
		str := StrApply(funcTypeExemplar, s.Cfg, string(cardState), string(cardInfo), string(devInfo))
		fmt.Println("***", str)
		return RateTripResult(str)
	}
}

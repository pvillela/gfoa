package fn

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
)

// ValidateTripRequestBfS contains the dependencies required for the construction of a
// ValidateTripRequestBf.
type ValidateTripRequestBfS struct {
	Cfg string
}

// ValidateTripRequestBf is the type of a funtion that validates a trip request.
type ValidateTripRequestBf = func(rpc.TripRequest)

// Make creates a ValidateTripRequestBf from its dependencies.
func (s ValidateTripRequestBfS) Make() ValidateTripRequestBf {
	var funcTypeExemplar ValidateTripRequestBfS
	return func(tripRequest rpc.TripRequest) {
		str := StrApply(funcTypeExemplar, s.Cfg, tripRequest.ToString())
		fmt.Println("***", str)
	}
}

package fn

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/event"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/rpc"
)

// TripSvcflowS contains the dependencies required for the construction of a
// TripSvcflow.
type TripSvcflowS struct {
	ValidateTripRequestBf   func(rpc.TripRequest)
	GetUpdatedCardStateFlow func(model.CardInfo) model.CardState
	RateTripSc              func(model.CardState, model.CardInfo, model.DeviceInfo) RateTripResult
	PrepareUsageBf          func(model.CardInfo, model.DeviceInfo, model.CardState, RateTripResult) model.Usage
	WriteUsageDaf           func(model.Usage)
	PrepareDownstreamEvtBf  func(model.Usage) event.DownstreamEvt
	DownstreamEp            func(event.DownstreamEvt)
	PrepareResponseBf       func(model.CardState, RateTripResult) rpc.TripResponse
}

// TripSvcflow is the type of a funtion that receives a trip request from a device,
// rates the trip, updates the database, and returns a response that includes the decision
// on whether the trip is allowed and, if applicable, the fare charged.
type TripSvcflow = func(rpc.TripRequest) rpc.TripResponse

// Make creates a TripSvcflow from its dependencies.
func (s TripSvcflowS) Make() TripSvcflow {
	var funcTypeExemplar TripSvcflowS
	return func(tripRequest rpc.TripRequest) rpc.TripResponse {
		cardInfo := tripRequest.CardInfo
		deviceInfo := tripRequest.DeviceInfo
		s.ValidateTripRequestBf(tripRequest)
		cardState := s.GetUpdatedCardStateFlow(cardInfo)
		rateTripResult := s.RateTripSc(cardState, cardInfo, deviceInfo)
		usage := s.PrepareUsageBf(cardInfo, deviceInfo, cardState, rateTripResult)
		s.WriteUsageDaf(usage)
		evt := s.PrepareDownstreamEvtBf(usage)
		s.DownstreamEp(evt)
		tripResponse := s.PrepareResponseBf(cardState, rateTripResult)
		str := StrApply(funcTypeExemplar, "", tripRequest.ToString())
		fmt.Println("***", str)
		return tripResponse
	}
}

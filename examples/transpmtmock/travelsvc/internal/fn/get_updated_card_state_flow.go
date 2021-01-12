package fn

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"
)

// GetUpdatedCardStateFlowS contains the dependencies required for the construction of a
// GetUpdatedCardStateFlow.
type GetUpdatedCardStateFlowS struct {
	GetPostedCardStateDaf func(model.CardInfo) model.CardState
	GetUnpostedUsagesDaf  func(model.CardInfo) []model.Usage
	UpdateCardStateBf     func(model.CardState, []model.Usage) model.CardState
}

// GetUpdatedCardStateFlow retrieves information from the database and creates an updated
// card state in memory based on the posted state and unposted usages.
type GetUpdatedCardStateFlow = func(model.CardInfo) model.CardState

// Make creates a GetUpdatedCardStateFlow from its dependencies.
func (s GetUpdatedCardStateFlowS) Make() GetUpdatedCardStateFlow {
	var funcTypeExemplar GetUpdatedCardStateFlowS
	return func(cardInfo model.CardInfo) model.CardState {
		postedCardState := s.GetPostedCardStateDaf(cardInfo)
		unpostedUsages := s.GetUnpostedUsagesDaf(cardInfo)
		updatedCardState := s.UpdateCardStateBf(postedCardState, unpostedUsages)
		str := StrApply(funcTypeExemplar, "", string(cardInfo))
		fmt.Println("***", str)
		return updatedCardState
	}
}

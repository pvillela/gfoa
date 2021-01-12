package fn

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"
)

// GetPostedCardStateDafS contains the dependencies required for the construction of a
// GetPostedCardStateDaf.
type GetPostedCardStateDafS struct {
	Cfg string
}

// GetPostedCardStateDaf is the type of a funtion that retrieves the latest posted card state
// for a given card from the database.
type GetPostedCardStateDaf = func(model.CardInfo) model.CardState

// Make creates a GetPostedCardStateDaf from its dependencies.
func (s GetPostedCardStateDafS) Make() GetPostedCardStateDaf {
	var funcTypeExemplar GetPostedCardStateDafS
	return func(cardInfo model.CardInfo) model.CardState {
		str := StrApply(funcTypeExemplar, s.Cfg, string(cardInfo))
		fmt.Println("***", str)
		return model.CardState(str)
	}
}

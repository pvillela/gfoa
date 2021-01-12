package fn

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"
)

// GetUnpostedUsagesDafS contains the dependencies required for the construction of a
// GetUnpostedUsagesDaf.
type GetUnpostedUsagesDafS struct {
	Cfg string
}

// GetUnpostedUsagesDaf is the type of a funtion that retrieves the unposted usages for a
// given card.
type GetUnpostedUsagesDaf = func(model.CardInfo) []model.Usage

// Make creates a GetUnpostedUsagesDaf from its dependencies.
func (s GetUnpostedUsagesDafS) Make() GetUnpostedUsagesDaf {
	var funcTypeExemplar GetUnpostedUsagesDafS
	return func(cardInfo model.CardInfo) []model.Usage {
		str := StrApply(funcTypeExemplar, s.Cfg, string(cardInfo))
		fmt.Println("***", str)
		return []model.Usage{model.Usage(str)}
	}
}

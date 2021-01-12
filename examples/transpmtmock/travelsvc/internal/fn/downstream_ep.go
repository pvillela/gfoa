package fn

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/event"
)

// DownstreamEpS contains the dependencies required for the construction of a
// DownstreamEp.
type DownstreamEpS struct {
	Cfg string
}

// DownstreamEp is the type of a funtion that publishes events for downstream processing.
type DownstreamEp = func(event.DownstreamEvt)

// Make creates a DownstreamEp from its dependencies.
func (s DownstreamEpS) Make() DownstreamEp {
	var funcTypeExemplar DownstreamEpS
	return func(evt event.DownstreamEvt) {
		str := StrApply(funcTypeExemplar, s.Cfg, string(evt))
		fmt.Println("***", str)
	}
}

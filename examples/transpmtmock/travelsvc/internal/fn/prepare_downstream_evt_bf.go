package fn

import (
	"fmt"

	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/event"
	"github.com/pvillela/gfoa/examples/transpmtmock/travelsvc/pkg/model"
)

// PrepareDownstreamEvtBfS contains the dependencies required for the construction of a
// PrepareDownstreamEvtBf.
type PrepareDownstreamEvtBfS struct {
	Cfg string
}

// PrepareDownstreamEvtBf is the type of a funtion that creates a DownstreamEvt from a usage.
type PrepareDownstreamEvtBf = func(model.Usage) event.DownstreamEvt

// Make creates a PrepareDownstreamEvtBf from its dependencies.
func (s PrepareDownstreamEvtBfS) Make() PrepareDownstreamEvtBf {
	var funcTypeExemplar PrepareDownstreamEvtBfS
	return func(usage model.Usage) event.DownstreamEvt {
		str := StrApply(funcTypeExemplar, s.Cfg, string(usage))
		fmt.Println("***", str)
		return event.DownstreamEvt(str)
	}
}

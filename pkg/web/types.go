package web

type Filler = func(interface{}) error

type HandlerHelper = func(Filler) interface{}

type WithHandlerHelper interface {
	HandlerHelper() HandlerHelper
}

type PseudoPostHandler = func(Filler) (interface{}, error)

type Any = interface{}

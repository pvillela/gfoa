package web

type Filler = func(interface{})

type HandlerHelper = func(Filler) interface{}

type WithHandlerHelper interface {
	HandlerHelper() HandlerHelper
}

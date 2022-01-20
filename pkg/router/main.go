package router

import (
	impl "github.com/umatare5/logbook-api-router-impl"
)

// RegisterHandlers ...
func RegisterHandlers(r impl.EchoRouter, si impl.ServerInterface) {
	impl.RegisterHandlers(r, si)
}

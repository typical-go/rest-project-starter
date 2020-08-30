package server

import (
	"github.com/typical-go/typical-rest-server/pkg/typrest"
	"go.uber.org/dig"
)

type (
	// Router for server
	Router struct {
		dig.In
	}
)

var _ (typrest.Router) = (*Router)(nil)

// SetRoute set route
func (*Router) SetRoute(e typrest.Server) error {
	routers := typrest.Routers{
		// TODO: add controller here
		// example at https://github.com/typical-go/typical-rest-server/blob/master/internal/app/server/router.go
	}
	return routers.SetRoute(e)
}

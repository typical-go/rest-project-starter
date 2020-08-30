package profiler

import (
	"github.com/labstack/echo/v4"
	"github.com/typical-go/typical-rest-server/pkg/typrest"
	"go.uber.org/dig"
)

type (
	// HealthCheck health check
	HealthCheck struct {
		dig.In
	}
)

var _ typrest.Router = (*HealthCheck)(nil)

// SetRoute to profiler api
func (h *HealthCheck) SetRoute(e typrest.Server) error {
	e.Any("application/health", h.handle)
	return nil
}

func (h *HealthCheck) handle(ec echo.Context) error {
	hc := typrest.HealthCheck{
		// TODO: add health function
		// example at https://github.com/typical-go/typical-rest-server/blob/master/internal/app/profiler/health_check.go
	}
	status, message := hc.Result()
	return ec.JSON(status, message)
}

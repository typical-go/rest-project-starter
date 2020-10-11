package infra

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-rest-server/pkg/typrest"
	"go.uber.org/dig"
)

type (
	// HealthCheck health check
	HealthCheck struct {
		dig.In
	}
)

// Handle health check
func (h *HealthCheck) Handle(ec echo.Context) error {
	healthy, detail := typrest.HealthStatus(typrest.HealthMap{
		// TODO:
	})

	status := http.StatusOK
	if !healthy {
		status = http.StatusServiceUnavailable
	}

	return ec.JSON(status, map[string]interface{}{
		"name":   fmt.Sprintf("%s (%s)", typgo.ProjectName, typgo.ProjectVersion),
		"status": detail,
	})
}

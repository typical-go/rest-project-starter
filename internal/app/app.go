package app

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/typical-go/rest-project-starter/internal/app/infra"
	"github.com/typical-go/rest-project-starter/internal/app/infra/log"
	"github.com/typical-go/typical-rest-server/pkg/logruskit"
	"github.com/typical-go/typical-rest-server/pkg/typrest"
	"go.uber.org/dig"

	// enable `/debug/vars`
	_ "expvar"

	// enable `/debug/pprof` API
	_ "net/http/pprof"
)

const (
	healthCheckPath = "/application/health"
)

type (
	app struct {
		dig.In
		Config      *infra.AppCfg
		HealthCheck infra.HealthCheck
	}
)

// Start app
func Start(a app) (err error) {
	e := echo.New()
	defer shutdown(e)

	e.HideBanner = true
	e.Debug = a.Config.Debug

	logger := log.SetDebug(a.Config.Debug)
	e.Logger = logruskit.EchoLogger(logger)

	setMiddleware(a, e)
	setRoute(a, e)

	return e.StartServer(&http.Server{
		Addr:         a.Config.Address,
		ReadTimeout:  a.Config.ReadTimeout,
		WriteTimeout: a.Config.WriteTimeout,
	})
}

func setMiddleware(a app, e *echo.Echo) {
	e.Use(log.Middleware)
	e.Use(middleware.Recover())
}

func setRoute(a app, e *echo.Echo) {
	// typrest.SetRoute(e,
	// 	// TODO: add domain router
	// )

	e.GET(healthCheckPath, a.HealthCheck.Handle)
	e.HEAD(healthCheckPath, a.HealthCheck.Handle)
	e.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))
	e.GET("/debug/*/*", echo.WrapHandler(http.DefaultServeMux))

	if a.Config.Debug {
		logrus.Debugf("Print routes:\n  %s\n\n",
			strings.Join(typrest.DumpEcho(e), "\n  "))
	}
}

func shutdown(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	e.Shutdown(ctx)
}

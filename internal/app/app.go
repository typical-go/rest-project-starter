package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/typical-go/rest-project-starter/internal/app/infra"
	"github.com/typical-go/rest-project-starter/internal/app/profiler"
	"github.com/typical-go/rest-project-starter/internal/app/server"
	"github.com/typical-go/typical-rest-server/pkg/typrest"
	"go.uber.org/dig"
)

type (
	app struct {
		dig.In
		Config   *infra.AppCfg
		Profiler profiler.Router
		Server   server.Router
	}
)

// Start app
func Start(a app) (err error) {
	e := echo.New()
	defer Shutdown(e)

	e.HideBanner = true
	e.Debug = a.Config.Debug

	a.SetLoggger(e)
	a.SetMiddleware(e)

	if err := a.SetRoute(e); err != nil {
		return err
	}
	return e.StartServer(&http.Server{
		Addr:         a.Config.Address,
		ReadTimeout:  a.Config.ReadTimeout,
		WriteTimeout: a.Config.WriteTimeout,
	})
}

// Shutdown app
func Shutdown(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	e.Shutdown(ctx)
}

// SetMiddleware set app middleware
func (a app) SetMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	if e.Debug {
		e.Use(loggerMiddleware())
	}
}

// SetRoute set app route
func (a app) SetRoute(e *echo.Echo) error {
	routers := typrest.Routers{
		&a.Server,
		&a.Profiler,
	}
	return routers.SetRoute(e)
}

// SetLogger set app logger
func (a app) SetLoggger(e *echo.Echo) {
	logger := logrus.StandardLogger()     // NOTE: always use standard logrus logger
	e.Logger = typrest.WrapLogrus(logger) // NOTE: setup echo logger
	log.SetOutput(logger.Writer())        // NOTE: std golang log use same output writer with logrus

	if a.Config.Debug {
		logger.SetLevel(logrus.DebugLevel)
		logger.SetFormatter(&logrus.TextFormatter{})
	} else {
		logger.SetLevel(logrus.WarnLevel)
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
}

// loggerMiddleware log every request
func loggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			bytesIn := req.Header.Get(echo.HeaderContentLength)

			logrus.WithFields(map[string]interface{}{
				"status":    res.Status,
				"latency":   stop.Sub(start).String(),
				"bytes_in":  bytesIn,
				"bytes_out": strconv.FormatInt(res.Size, 10),
			}).Info(fmt.Sprintf("%s %s", req.Method, req.RequestURI))
			return nil
		}
	}
}

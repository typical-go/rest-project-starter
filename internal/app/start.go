package app

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

// Start the app
func Start(a app) (err error) {
	e := echo.New()
	defer shutdown(e)

	e.HideBanner = true
	e.Debug = a.Debug

	a.SetLoggger(e)
	a.SetMiddleware(e)

	if err := a.SetRoute(e); err != nil {
		return err
	}
	return e.Start(a.Address)
}

func shutdown(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	e.Shutdown(ctx)
}

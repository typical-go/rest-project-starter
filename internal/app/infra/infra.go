package infra

import (
	"go.uber.org/dig"
)

type (
	// Infra infrastructure for the project
	Infra struct {
		// TODO: uncomment `dig.Out` and add infra
		// dig.Out
	}
	setupParam struct {
		dig.In
		// TODO: add param
	}
	teardownParam struct {
		dig.In
		// TODO: add param
	}
)

// Setup infra
// @ctor
func Setup(p setupParam) Infra {
	// TODO: add setup function e.g. create connection, etc
	// Use `logrus.Fatal` to handling error since infra component is mandatory
	return Infra{}
}

// Teardown infra
// @dtor
func Teardown(p teardownParam) error {
	// TODO: add teardown function e.g. close database, etc.
	return nil
}

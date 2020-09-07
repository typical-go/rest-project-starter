package infra

import (
	"go.uber.org/dig"
)

type (
	// Infra infrastructure for the project
	Infra struct {
		dig.Out
		// TODO: add infra
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
func Setup(p setupParam) (infras Infra, err error) {
	// TODO: add setup function e.g. create connection, etc
	return Infra{}, nil
}

// Teardown infra
// @dtor
func Teardown(p teardownParam) error {
	// TODO: add teardown function e.g. close database, etc.
	return nil
}

package infra

type (
	// AppCfg application configuration
	// @envconfig (prefix:"APP")
	AppCfg struct {
		Address string `envconfig:"ADDRESS" default:":8089" required:"true"`
		Debug   bool   `envconfig:"DEBUG" default:"true"`
	}
)

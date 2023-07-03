package config

import (
	"github.com/bitxeno/go-docker-skeleton/internal/app"
	"github.com/creasty/defaults"
)

var Custom CustomConfiguration

type CustomConfiguration struct {
	Test string `koanf:"test" default:"test"`
}

func loadCustom() error {
	// set default value
	if err := defaults.Set(&Custom); err != nil {
		return err
	}

	// load from default config.yaml
	if err := app.Cfg().BindStruct("custom", &Custom); err != nil {
		return err
	}
	return nil
}

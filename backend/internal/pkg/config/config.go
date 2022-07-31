package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/mcuadros/go-defaults"
)

//config contains all global config of application
type config struct {
	Env        string           `env:"ENV"`
	Http       HttpConfig       `envPrefix:"HTTP_"`
	DB         DBConfig         `envPrefix:"DB_"`
	Blockchain BlockchainConfig `envPrefix:"BLC_"`
}

var cfg *config

//Instance return instance of global config
func Instance() *config {
	if cfg == nil {
		cfg = &config{
			Http:       HttpConfig{},
			DB:         DBConfig{},
			Blockchain: BlockchainConfig{},
		}
	}
	return cfg
}

//Load loads configurations from file and env
func Load(configFile string) error {
	// Default config values
	c := Instance()
	defaults.SetDefaults(c)

	if err := env.Parse(c); err != nil {
		return err
	}

	return nil
}

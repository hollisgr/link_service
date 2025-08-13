package config

import (
	"link_service/pkg/logger"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Addr      string
	BindIP    string `env:"BIND_IP" env-default:"127.0.0.1"`
	Port      string `env:"LISTEN_PORT" env-default:"8080"`
	FileTypes string `env:"FILE_TYPES" env-default:"image/jpeg,application/pdf"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *logger.Logger) *Config {

	once.Do(func() {
		// logger.Infoln("reading app configuration")
		instance = &Config{}
		err := cleanenv.ReadConfig(".env", instance)
		if err != nil {
			// logger.Infoln("read app configuration error, using default settings")
		} else {
			// logger.Infoln("config OK")
		}
		instance.Addr = instance.BindIP + ":" + instance.Port
	})
	return instance
}

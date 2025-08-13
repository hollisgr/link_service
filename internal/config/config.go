package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Addr      string
	BindIP    string `env:"BIND_IP"`
	Port      string `env:"LISTEN_PORT"`
	FileTypes string `env:"FILE_TYPES"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {

	once.Do(func() {
		log.Println("reading app configuration")
		instance = &Config{}
		err := cleanenv.ReadConfig(".env", instance)
		if err != nil {
			log.Fatalln("read app configuration error")
		} else {
			log.Println("config OK")
		}
		instance.Addr = instance.BindIP + ":" + instance.Port
	})
	return instance
}

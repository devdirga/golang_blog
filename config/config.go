package config

import (
	"os"

	"github.com/tkanos/gonfig"
)

type Conf struct {
	DB            string
	Redis         string
	RedisPassword string
	IsDebug       bool
	IsConcurrent  bool
	Secret        string
	GoogleSmtpKey string
	URLFront      string
}

var conf Conf

func Init() {
	err := gonfig.GetConf("config.json", &conf)
	if err != nil {
		os.Exit(500)
	}
}

func GetConf() Conf {
	return conf
}

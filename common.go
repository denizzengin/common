package common 

import (
	"github.com/denizzengin/config"
)

func GetEnv() string {
	return config.Config.Environment.Env
}

func IsDevelopment() bool {
	return config.Config.Environment.Env == Development
}

func IsProduction()bool{
	return config.Config.Environment.Env == Production
}

func IsTest() bool {
	return config.Config.Environment.Env == Test
}

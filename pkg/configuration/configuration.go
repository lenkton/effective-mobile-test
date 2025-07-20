package configuration

import (
	"fmt"
	"os"
)

type Configuration struct {
	port string
}

func New() Configuration {
	return Configuration{
		port: fetchEnv("PORT", "8080"),
	}
}

func (config Configuration) Address() string {
	return fmt.Sprintf(":%v", config.port)
}

func fetchEnv(name string, defaultValue string) string {
	value, found := os.LookupEnv(name)
	if found {
		return value
	} else {
		return defaultValue
	}
}

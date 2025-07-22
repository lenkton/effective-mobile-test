package configuration

import (
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	Port        string
	DatabaseURL string
}

func New() Configuration {
	return Configuration{
		Port:        fetchEnv("PORT", "8080"),
		DatabaseURL: mustFetchEnv("DB_URL"),
	}
}

func (config Configuration) Address() string {
	return fmt.Sprintf(":%v", config.Port)
}

func fetchEnv(name string, defaultValue string) string {
	value, found := os.LookupEnv(name)
	if found {
		return value
	} else {
		return defaultValue
	}
}

func mustFetchEnv(name string) string {
	value, found := os.LookupEnv(name)
	if found {
		return value
	} else {
		log.Fatalf("Missing environment variable: %s\n", name)
		return ""
	}
}

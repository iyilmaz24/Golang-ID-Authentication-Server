package config

import (
	"log"
	"os"
	"strings"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	DSN  string
	Port string
	Cors map[string]bool

	// CertFile string
	// KeyFile string
}

func LoadConfig() *Config {
	once.Do(func() { // ensure that the config is only loaded once

		dsn, ok := os.LookupEnv("DB_DSN")
		if !ok {
			log.Fatal("DB_DSN is not set in environment variables")
		}
		port := ":8200"

		corsString := os.Getenv("CORS_ORIGIN")
		if corsString == "" {
			log.Fatal("$CORS_ORIGIN env variable not set")
		}
		corsUrls := strings.Split(corsString, ",")

		corsOrigin := make(map[string]bool)
		for _, url := range corsUrls {
			corsOrigin[url] = true
		}

		// certFile := os.Getenv("CERT_PATH")
		// keyFile := os.Getenv("KEY_PATH")

		// if certFile == "" || keyFile == "" {
		// 	log.Fatal("CERT_PATH and KEY_PATH environment variables must be set for HTTPS")
		// }

		instance = &Config{
			DSN:  dsn,
			Port: port,
			Cors: corsOrigin,

			// CertFile: certFile,
			// KeyFile: keyFile,
		}

	})

	return instance
}
package config

import (
	"log"
	"os"
)

type Config struct {
	DSN  string
	Port string
	// CertFile string
	// KeyFile string
}

func LoadConfig() *Config {
	dsn, ok := os.LookupEnv("DB_DSN")
	if !ok {
		log.Fatal("DB_DSN is not set in environment variables")
	}

	port := ":8200"

	// port, ok := os.LookupEnv("PORT")
	// if !ok {
	// 	port = ":8200"
	// }

	// certFile := os.Getenv("CERT_PATH")
    // keyFile := os.Getenv("KEY_PATH")

	// if certFile == "" || keyFile == "" {
	// 	log.Fatal("CERT_PATH and KEY_PATH environment variables must be set for HTTPS")
	// }

	return &Config{
		DSN:  dsn,
		Port: port,
		// CertFile: certFile,
		// KeyFile: keyFile,
	}
}
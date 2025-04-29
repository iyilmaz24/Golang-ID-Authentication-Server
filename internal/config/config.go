// Package config provides configuration management using AWS Systems Manager Parameter Store
package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var (
	once     sync.Once
	instance *Config
)

type Config struct {
	DbDsn         string
	Port          string
	Cors          map[string]bool
	AdminPassword string
	// CertFile string
	// KeyFile string
}

type ConfigDefinition struct {
	Path         string
	Type         string
	DefaultValue string
}

var configDefinitions = map[string]ConfigDefinition{
	"CORS_ORIGIN": {
		Path: "/backend/internal/id-auth-cors-origin",
		Type: "StringList",
	},
	"DB_DSN": {
		Path: "/backend/internal/db_dsn",
		Type: "SecureString",
	},
	"PORT": {
		Path:         "/backend/ports/id-auth",
		Type:         "String",
		DefaultValue: ":8200",
	},
	"ADMIN_PASSWORD": {
		Path: "/backend/internal/admin-password",
		Type: "SecureString",
	},
}

func getSystemsManagerParameter(paramName string, ssmClient *ssm.Client) string {

	paramInfo, exists := configDefinitions[paramName]
	if !exists {
		log.Fatalf("***ERROR (config): Parameter '%s' not found in configDefinitions", paramName)
	}
	isEncrypted := paramInfo.Type == "SecureString"

	log.Printf("Attempting to retrieve parameter: %s (Path: %s)", paramName, paramInfo.Path)

	param, err := ssmClient.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(paramInfo.Path),
		WithDecryption: aws.Bool(isEncrypted),
	})

	if err != nil {
		log.Printf("ERROR retrieving parameter %s: %v", paramName, err)

		username, _ := os.Hostname()
		log.Printf("Hostname: %s", username)

		if paramInfo.DefaultValue != "" {
			log.Printf("Using default value for %s", paramName)
			return paramInfo.DefaultValue
		}
		errorMsg := fmt.Sprintf("***ERROR (config): Failed to retrieve parameter '%s' from Systems Manager: %v", paramName, err)
		log.Fatal(errorMsg)
	}
	log.Printf("Successfully retrieved parameter: %s", paramName)

	return *param.Parameter.Value
}

func LoadConfig() *Config {
	once.Do(func() {

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion("us-east-1"), // Specify your AWS region
		)
		if err != nil {
			log.Fatal("***ERROR (config): Unable to load AWS SDK config: ", err)
		}
		log.Println("AWS SDK Config loaded successfully")

		ssmClient := ssm.NewFromConfig(cfg)

		corsString := getSystemsManagerParameter("CORS_ORIGIN", ssmClient)
		corsUrls := strings.Split(corsString, ",")

		corsOrigin := make(map[string]bool, len(corsUrls))
		for _, url := range corsUrls {
			trimmedURL := strings.TrimSpace(url)
			if trimmedURL != "" {
				corsOrigin[trimmedURL] = true
			}
		}

		port := getSystemsManagerParameter("PORT", ssmClient)
		dbDsn := getSystemsManagerParameter("DB_DSN", ssmClient)
		adminPassword := getSystemsManagerParameter("ADMIN_PASSWORD", ssmClient)

		instance = &Config{
			Port:          port,
			Cors:          corsOrigin,
			DbDsn:         dbDsn,
			AdminPassword: adminPassword,
		}

		log.Println("Configuration loaded successfully")
	})

	return instance
}

package config

import (
	"log"
	"os"
)

type AppConfig struct {
	Port                string
	KeycloakBaseURL     string
	KeycloakRealm       string
	KeycloakClientID    string
	KeycloakRedirectURI string
}

func Load() AppConfig {
	config := AppConfig{
		Port:                getEnv("PORT", "8080"),
		KeycloakBaseURL:     mustGet("KEYCLOAK_BASE_URL"),
		KeycloakRealm:       mustGet("KEYCLOAK_REALM"),
		KeycloakClientID:    mustGet("KEYCLOAK_CLIENT_ID"),
		KeycloakRedirectURI: mustGet("KEYCLOAK_REDIRECT_URI"),
	}

	return config
}

func mustGet(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Variável obrigatória não definida: %s", key)
	}
	return value
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

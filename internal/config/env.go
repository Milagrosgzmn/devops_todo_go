package config

import "os"

// GetEnv obtiene una variable de entorno o retorna un valor por defecto
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetStringOrDefault(key, defaultValue string) string {
	val := viper.GetString(key)
	if val != "" {
		return val
	}

	return defaultValue
}

func GetBoolOrDefault(key string, defaultValue bool) bool {
	if exists := viper.IsSet(key); exists {
		return viper.GetBool(key)
	}
	return defaultValue
}

func GetIntOrDefault(key string, defaultValue int) int {
	if exists := viper.IsSet(key); exists {
		return viper.GetInt(key)
	}
	return defaultValue
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetStringNotEmpty(key string) (string, error) {
	val := viper.GetString(key)
	if val != "" {
		return val, nil
	}

	return "", fmt.Errorf("configuration required: %s", key)
}

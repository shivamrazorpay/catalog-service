package common

import (
	"github.com/matoous/go-nanoid/v2"
	"os"
)

// UniqueId generates a unique id
func UniqueId() (string, error) {
	id, err := gonanoid.New(10)
	if err != nil {
		return "", err
	}
	return id, nil
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

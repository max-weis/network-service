package config

import (
	"os"
)

type Config struct {
	RedisHost  string
	RedisPort  string
	NmapTarget string
}

func NewConfig() Config {
	return Config{
		RedisHost:  getEnvVar("REDIS_HOST", "localhost"),
		RedisPort:  getEnvVar("REDIS_PORT", "6379"),
		NmapTarget: getEnvVar("TARGET", "192.168.178.1/24"),
	}
}

func getEnvVar(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		val = def
	}
	return val
}

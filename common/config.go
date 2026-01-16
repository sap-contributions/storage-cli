package common

import (
	"log/slog"
	"sync"
)

type Config struct {
	LogLevel slog.Level
}

var (
	instance *Config
	once     sync.Once
)

func InitConfig(logLevel slog.Level) {
	once.Do(func() {
		instance = &Config{LogLevel: logLevel}
	})
}

func GetConfig() *Config {
	return instance
}

func IsDebug() bool {
	if instance == nil {
		return false
	}
	if instance.LogLevel == slog.LevelDebug {
		return true
	}
	return false
}

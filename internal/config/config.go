package config

import (
	"errors"
	"ohmyies/pkg/filestore"
	"ohmyies/pkg/logger"
	"os"
)

type Config struct {
	LogFile string `json:"log_file"` // file where logs will be written
	Debug   bool   `json:"debug"`    // write debug logs
}

const (
	defaultLogFile = "app.log"
	defaultDebug   = false
)

func MustLoadOrCreateConfig(configFile string) *Config {
	deviceConfig, err := LoadOrCreateConfig(configFile)
	if err != nil {
		logger.Fatalf("MustLoadOrCreateConfig: failed to load/create config: %v", err)
		return nil
	}
	return deviceConfig
}

func LoadOrCreateConfig(configFile string) (*Config, error) {
	var cfg *Config
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		cfg = &Config{
			LogFile: defaultLogFile,
			Debug:   defaultDebug,
		}
		return cfg, filestore.SaveJSON(configFile, cfg)
	} else {
		if err = filestore.LoadJSON(configFile, &cfg); err != nil {
			return nil, err
		}
		return cfg, nil
	}
}

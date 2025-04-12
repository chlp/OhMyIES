package config

import (
	"errors"
	"ohmyies/internal/model"
	"ohmyies/pkg/filestore"
	"ohmyies/pkg/logger"
	"os"
)

type Config struct {
	LogFile string       `json:"log_file"` // file where logs will be written
	Debug   bool         `json:"debug"`    // write debug logs
	Feeds   []FeedConfig `json:"feeds"`    // list of feeds to fetch
}

type FeedConfig struct {
	Name string       `json:"name"`
	Key  string       `json:"key"`
	Key2 string       `json:"key_2"`
	Chat []ChatConfig `json:"chat"`
}

type ChatConfig struct {
	Type        model.MsgType `json:"type"`
	BotApiToken string        `json:"bot_api_token"`
	ChatId      string        `json:"chat_id"`
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

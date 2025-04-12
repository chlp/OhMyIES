package main

import (
	"fmt"
	"ohmyies/internal/config"
	"ohmyies/internal/model"
	"ohmyies/internal/rss"
	"ohmyies/internal/tg"
	"ohmyies/pkg/application"
	"ohmyies/pkg/filestore"
	"ohmyies/pkg/logger"
	"os"
)

const (
	appName           = "ohmyies"
	defaultConfigFile = "config.json"
)

func main() {
	configPath := defaultConfigFile
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}
	cfg := config.MustLoadOrCreateConfig(configPath)
	logger.InitLogger(appName, cfg.LogFile, cfg.Debug)
	_, appDone := application.NewApp(appName, cfg.LogFile, cfg.Debug)

	for _, feedCfg := range cfg.Feeds {
		chatsByType := make(map[model.MsgType]*tg.Chat)
		for _, chatCfg := range feedCfg.Chat {
			chat := tg.NewChat(chatCfg.BotApiToken, chatCfg.ChatId)
			chatsByType[chatCfg.Type] = chat
		}
		feedStore := filestore.NewFileStore("feed_guids-" + feedCfg.Name + ".json")
		rss.NewFeed(feedCfg.Name, feedCfg.Key, feedCfg.Key2, feedStore, func(msg model.Msg) bool {
			if chat, ok := chatsByType[msg.Type]; ok {
				return chat.SendMessage(fmt.Sprintf("%s\n*%s*\n\n%s",
					msg.PubDate.Format("2006-01-02 15:04:05"), msg.Title, msg.Description))
			}
			return true
		})
	}

	<-appDone
	logger.Printf("Application gracefully shut down")
}

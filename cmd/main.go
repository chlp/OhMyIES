package main

import (
	"ohmyies/internal/config"
	"ohmyies/internal/store"
	"ohmyies/internal/worker"
	"ohmyies/pkg/application"
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

	app, appDone := application.NewApp(appName, cfg.LogFile, cfg.Debug)

	ohMyIesConfig := config.LoadApiConf()

	fileStore := store.NewFileStore(ohMyIesConfig.DB)
	runnerStore := store.NewRunnerStore()
	taskStore := store.NewTaskStore()

	worker.NewWorker(appCtx, fileStore, runnerStore).Run()

	<-appDone
	logger.Printf("Application gracefully shut down")
}

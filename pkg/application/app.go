package application

import (
	"context"
	"ohmyies/pkg/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type App struct {
	Ctx context.Context
	Wg  sync.WaitGroup
}

func NewApp(name, logFile string, debug bool) (*App, <-chan struct{}) {
	logger.InitLogger(name, logFile, debug)

	appCtx, appShutdown := context.WithCancel(context.Background())
	app := &App{
		Ctx: appCtx,
		Wg:  sync.WaitGroup{},
	}
	appDone := make(chan struct{})

	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
		<-done

		appShutdown()
		app.Wg.Wait()
		close(appDone)
	}()

	logger.Printf("App::NewApp")

	return app, appDone
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/atmxlab/proxychecker/cmd/app"
	"github.com/atmxlab/proxychecker/internal/api/proxychecker"
	"github.com/atmxlab/proxychecker/internal/pkg/config"
	_ "github.com/atmxlab/proxychecker/pkg/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	appCfg, err := config.LoadAndScan[app.Config]("config/api/api.atmc")
	if err != nil {
		panic(err)
	}

	setupLogger(appCfg.Logger.Level)

	baseApp := app.NewApp(app.SetupContainerBuilder(appCfg).Build())

	a := NewApp(baseApp, proxychecker.New(baseApp.Container()))

	a.Init()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = a.Start(ctx); err != nil {
			logrus.Errorf("app.Start error: %s", err)
		}
	}()

	<-ctx.Done()
	if err = a.Stop(ctx); err != nil {
		logrus.Errorf("app.Stop error: %s", err)
	}
	wg.Wait()
}

func setupLogger(level string) {
	m := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
	}
	if ll, ok := m[level]; ok {
		logrus.SetLevel(ll)
		logrus.Infof("Setup logger: level: [%s]", level)
	} else {
		panic(fmt.Errorf("invalid log level: %s", level))
	}
}

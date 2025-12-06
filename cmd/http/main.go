package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/atmxlab/proxychecker/cmd/app"
	"github.com/atmxlab/proxychecker/internal/pkg/config"
	_ "github.com/atmxlab/proxychecker/pkg/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	appCfg, err := config.LoadAndScan[app.Config]("config/app.atmc")
	if err != nil {
		panic(err)
	}

	cb := app.SetupContainerBuilder(appCfg)
	a := app.NewApp(cb.Build())

	a.Init()

	if err = a.Start(ctx); err != nil {
		panic(err)
	}
}

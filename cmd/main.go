package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/atmxlab/atmc"
	"github.com/atmxlab/proxychecker/cmd/app"
	_ "github.com/atmxlab/proxychecker/pkg/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	scanner, err := atmc.New().Load("config/config.atmc")
	if err != nil {
		panic(err)
	}

	var cfg app.Config
	if err = scanner.Scan(&cfg); err != nil {
		panic(err)
	}

	a := app.NewApp(cfg)

	a.Init()

	if err = a.Start(ctx); err != nil {
		panic(err)
	}
}

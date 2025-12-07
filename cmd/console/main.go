package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atmxlab/proxychecker/cmd/app"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/pkg/config"
	"github.com/atmxlab/proxychecker/internal/service/command"
	_ "github.com/atmxlab/proxychecker/pkg/logger"
	"github.com/atmxlab/proxychecker/pkg/waiter"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Checkers []string `atmc:"checkers"`
	Proxies  []string `atmc:"proxies"`
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	consoleCfg, err := config.LoadAndScan[Config]("config/console/console.atmc")
	if err != nil {
		panic(err)
	}
	appCfg, err := config.LoadAndScan[app.Config]("config/common/common.atmc")
	if err != nil {
		panic(err)
	}

	cb := app.SetupContainerBuilder(appCfg)
	a := app.NewApp(cb.Build())

	a.Init()

	go func() {
		if err = a.Start(ctx); err != nil {
			panic(err)
		}
	}()

	cmd := command.CheckInput{
		OperationTime: a.Container().Entities().TimeProvider().CurrentTime(ctx),
		Proxies:       consoleCfg.Proxies,
		Checkers: lo.Map(consoleCfg.Checkers, func(kind string, _ int) checker.Kind {
			return checker.KindFromString(kind)
		}),
	}

	res, err := a.Container().Commands().Check().Execute(ctx, cmd)
	if err != nil {
		panic(err)
	}

	logrus.Infof("check command result: taskGroupID: [%s]", res.TaskGroupID)

	if err = a.WaitTasksTerminated(ctx, waiter.WithTimeout(5*time.Minute)); err != nil {
		panic(err)
	}

	tasks, err := a.Container().Ports().GetTaskAggsByGroupID().Execute(ctx, res.TaskGroupID)
	if err != nil {
		panic(err)
	}

	printResult := func(tk *aggregate.Task) {
		fmt.Print("\n\n*********************************************\n")
		fmt.Printf("id: [%s]\n", tk.Task().ID())
		fmt.Printf("proxy: [%s]\n", tk.Proxy().URL())
		fmt.Printf("status: [%s]\n", tk.Task().Status())
		fmt.Printf("kind: [%s]\n", tk.Task().CheckerKind())
		fmt.Printf("result: [%s]\n", tk.Task().State().Result().String())
	}

	countryCount := make(map[string]int)
	failedTasks := 0
	for _, tk := range tasks {
		if tk.Task().Status() != task.StatusSuccess {
			printResult(tk)
			failedTasks++
		}
	}

	for _, tk := range tasks {
		if tk.Task().Status() == task.StatusSuccess {
			printResult(tk)
			if tk.Task().CheckerKind() == checker.KindGEO {
				countryCount[tk.Task().State().Result().GEOResult.Country]++
			}
		}
	}

	// TODO: grpc + http + swagger
	// TODO: contract with all info:
	//  - proxies count
	//  - tasks count (all + proxy)
	//  - success tasks count (all + proxy)
	//  - failure tasks count (all + proxy)
	//  - pending tasks count (all + proxy)
	//  - progress percentage
	//  - geo aggregate
	//  - protocol aggregate
	//  - proxy type aggregate
	//  - sort by success tasks and priority results
	//  - flag of done

	// TODO: proxy agg
	// TODO: group agg
	// TODO: group result query
	// TODO: deploy with docker
	fmt.Printf("all tasks: [%d]\n", len(tasks))
	fmt.Printf("success tasks: [%d]\n", len(tasks)-failedTasks)
	fmt.Printf("failed tasks: [%d]\n", failedTasks)
	fmt.Printf("valid proxies: [%d]\n", (len(tasks)-failedTasks)/2)

	for country, count := range countryCount {
		fmt.Printf("country [%s]: [%d]\n", country, count)
	}
}

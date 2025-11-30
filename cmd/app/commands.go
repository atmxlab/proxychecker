package app

import "github.com/atmxlab/proxychecker/internal/service/command"

type Commands struct {
	checkCommand *command.CheckCommand
}

func (c Commands) Check() *command.CheckCommand {
	return c.checkCommand
}

func (a *App) initCommands() {
	a.commands.checkCommand = command.NewCheckCommand(
		a.ports.runTx,
		a.ports.insertProxy,
		a.ports.insertTask,
		a.ports.scheduleTask,
	)
}

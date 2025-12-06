package app

import "github.com/atmxlab/proxychecker/internal/service/command"

type Commands struct {
	checkCommand *command.CheckCommand
}

func (c Commands) Check() *command.CheckCommand {
	return c.checkCommand
}

type CommandsBuilder struct {
	c *Container
}

func newCommandsBuilder(c *Container) *CommandsBuilder {
	return &CommandsBuilder{c: c}
}

func (cb *CommandsBuilder) Container() *Container {
	return cb.c
}

func (cb *CommandsBuilder) Check(c *command.CheckCommand) *CommandsBuilder {
	cb.c.commands.checkCommand = c
	return cb
}

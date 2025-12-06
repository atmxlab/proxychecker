package app

type Container struct {
	config   Config
	entities Entities
	ports    Ports
	commands Commands
	checkers Checkers
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Config() Config {
	return c.config
}

func (c *Container) Entities() Entities {
	return c.entities
}

func (c *Container) Ports() Ports {
	return c.ports
}

func (c *Container) Commands() Commands {
	return c.commands
}

func (c *Container) Checkers() Checkers {
	return c.checkers
}

type ContainerBuilder struct {
	c               *Container
	entitiesBuilder *EntitiesBuilder
	entitiesHooks   []func()
	commandsBuilder *CommandsBuilder
	commandsHooks   []func()
	checkersBuilder *CheckersBuilder
	checkersHooks   []func()
	portsBuilder    *PortsBuilder
	portsHooks      []func()
}

func NewContainerBuilder() *ContainerBuilder {
	cb := &ContainerBuilder{}
	cb.c = NewContainer()
	cb.portsBuilder = newPortsBuilder(cb.c)
	cb.checkersBuilder = newCheckersBuilder(cb.c)
	cb.entitiesBuilder = newEntitiesBuilder(cb.c)
	cb.commandsBuilder = newCommandsBuilder(cb.c)

	return cb
}

func (c *ContainerBuilder) WithConfig(config Config) *ContainerBuilder {
	c.entitiesHooks = append(c.entitiesHooks, func() {
		c.c.config = config
	})

	return c
}

func (c *ContainerBuilder) WithEntities(hook func(cb *EntitiesBuilder)) *ContainerBuilder {
	c.entitiesHooks = append(c.entitiesHooks, func() {
		hook(c.entitiesBuilder)
	})

	return c
}

func (c *ContainerBuilder) WithPorts(hook func(pb *PortsBuilder)) *ContainerBuilder {
	c.portsHooks = append(c.portsHooks, func() {
		hook(c.portsBuilder)
	})

	return c
}

func (c *ContainerBuilder) WithCheckers(hook func(cb *CheckersBuilder)) *ContainerBuilder {
	c.checkersHooks = append(c.checkersHooks, func() {
		hook(c.checkersBuilder)
	})

	return c
}

func (c *ContainerBuilder) WithCommands(hook func(cb *CommandsBuilder)) *ContainerBuilder {
	c.commandsHooks = append(c.commandsHooks, func() {
		hook(c.commandsBuilder)
	})

	return c
}

func (c *ContainerBuilder) Build() *Container {
	for _, hook := range c.entitiesHooks {
		hook()
	}
	for _, hook := range c.portsHooks {
		hook()
	}
	for _, hook := range c.checkersHooks {
		hook()
	}
	for _, hook := range c.commandsHooks {
		hook()
	}

	return c.c
}

package app

import (
	"github.com/atmxlab/proxychecker/internal/details/factory"
	"github.com/atmxlab/proxychecker/pkg/queue"
	"github.com/atmxlab/proxychecker/pkg/time"
)

type Entities struct {
	timeProvider   time.Provider
	clientFactory  *factory.ClientFactory
	ipApiFactory   *factory.IPApiFactory
	httpBinFactory *factory.HTTPBinFactory
	queue          *queue.Queue
}

func (e Entities) TimeProvider() time.Provider {
	return e.timeProvider
}

func (e Entities) ClientFactory() *factory.ClientFactory {
	return e.clientFactory
}

func (e Entities) IpApiFactory() *factory.IPApiFactory {
	return e.ipApiFactory
}

func (e Entities) HttpBinFactory() *factory.HTTPBinFactory {
	return e.httpBinFactory
}

func (e Entities) Queue() *queue.Queue {
	return e.queue
}

type EntitiesBuilder struct {
	c *Container
}

func newEntitiesBuilder(c *Container) *EntitiesBuilder {
	return &EntitiesBuilder{c: c}
}

func (eb *EntitiesBuilder) Container() *Container {
	return eb.c
}

func (eb *EntitiesBuilder) TimeProvider(tp time.Provider) *EntitiesBuilder {
	eb.c.entities.timeProvider = tp
	return eb
}

func (eb *EntitiesBuilder) ClientFactory(f *factory.ClientFactory) *EntitiesBuilder {
	eb.c.entities.clientFactory = f
	return eb
}

func (eb *EntitiesBuilder) IpAPIFactory(f *factory.IPApiFactory) *EntitiesBuilder {
	eb.c.entities.ipApiFactory = f
	return eb
}

func (eb *EntitiesBuilder) HTTPBinFactory(f *factory.HTTPBinFactory) *EntitiesBuilder {
	eb.c.entities.httpBinFactory = f
	return eb
}

func (eb *EntitiesBuilder) Queue(q *queue.Queue) *EntitiesBuilder {
	eb.c.entities.queue = q
	return eb
}

package app

import (
	"github.com/atmxlab/proxychecker/internal/service/task/handler"
)

type Checkers struct {
	geo        handler.Checker
	latency    handler.Checker
	externalIP handler.Checker
	url        handler.Checker
}

func (c Checkers) GEO() handler.Checker {
	return c.geo
}

func (c Checkers) Latency() handler.Checker {
	return c.latency
}

func (c Checkers) ExternalIP() handler.Checker {
	return c.externalIP
}

func (c Checkers) URL() handler.Checker {
	return c.url
}

type CheckersBuilder struct {
	c *Container
}

func newCheckersBuilder(c *Container) *CheckersBuilder {
	return &CheckersBuilder{c: c}
}

func (cb *CheckersBuilder) Container() *Container {
	return cb.c
}

func (cb *CheckersBuilder) GEO(g handler.Checker) *CheckersBuilder {
	cb.c.checkers.geo = g
	return cb
}

func (cb *CheckersBuilder) Latency(l handler.Checker) *CheckersBuilder {
	cb.c.checkers.latency = l
	return cb
}

func (cb *CheckersBuilder) ExternalIP(e handler.Checker) *CheckersBuilder {
	cb.c.checkers.externalIP = e
	return cb
}

func (cb *CheckersBuilder) URL(url handler.Checker) *CheckersBuilder {
	cb.c.checkers.url = url
	return cb
}

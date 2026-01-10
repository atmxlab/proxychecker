package app

import (
	"github.com/atmxlab/proxychecker/internal/service/task/handler"
)

type Checkers struct {
	geo        handler.Checker
	latency    handler.Checker
	externalIP handler.Checker
	url        handler.Checker
	https      handler.Checker
	mitm       handler.Checker
	proxyType  handler.Checker
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

func (c Checkers) HTTPS() handler.Checker {
	return c.https
}

func (c Checkers) MITM() handler.Checker {
	return c.mitm
}

func (c Checkers) Type() handler.Checker {
	return c.proxyType
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

func (cb *CheckersBuilder) HTTPS(https handler.Checker) *CheckersBuilder {
	cb.c.checkers.https = https
	return cb
}

func (cb *CheckersBuilder) MITM(mitm handler.Checker) *CheckersBuilder {
	cb.c.checkers.mitm = mitm
	return cb
}

func (cb *CheckersBuilder) Type(proxyType handler.Checker) *CheckersBuilder {
	cb.c.checkers.proxyType = proxyType
	return cb
}

package app

import (
	"github.com/atmxlab/proxychecker/internal/service/query"
)

type Queries struct {
	checkResult *query.CheckResultQuery
}

func (c Queries) CheckResult() *query.CheckResultQuery {
	return c.checkResult
}

type QueriesBuilder struct {
	c *Container
}

func newQueriesBuilder(c *Container) *QueriesBuilder {
	return &QueriesBuilder{c: c}
}

func (cb *QueriesBuilder) Container() *Container {
	return cb.c
}

func (cb *QueriesBuilder) CheckResult(c *query.CheckResultQuery) *QueriesBuilder {
	cb.c.queries.checkResult = c
	return cb
}

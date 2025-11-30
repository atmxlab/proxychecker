package acceptance_test

import (
	"context"
	"testing"
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/service/command"
	"github.com/atmxlab/proxychecker/internal/test"
	"github.com/stretchr/testify/require"
)

func TestCheckLatency(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
		now = time.Now()
	)
	app := test.NewApp(t)

	_, err := app.Commands().Check().Execute(
		ctx,
		command.CheckInput{
			OperationTime: now,
			Proxies:       []string{"https://proxy.io"},
			Checkers: []checker.Kind{
				checker.KindGEO,
			},
		},
	)
	require.NoError(t, err)
}

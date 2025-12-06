package latency_test

import (
	"context"
	"testing"
	stdtime "time"

	"github.com/atmxlab/proxychecker/internal/details/checker/latency"
	"github.com/atmxlab/proxychecker/internal/details/checker/latency/mocks"
	"github.com/atmxlab/proxychecker/internal/details/factory"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLatency(t *testing.T) {
	t.Parallel()

	var (
		ctx   = context.Background()
		tk    = entity.NewTaskBuilder().Build()
		px    = entity.NewProxyBuilder().Build()
		now   = stdtime.Now()
		since = 50 * stdtime.Millisecond
	)

	tp := mocks.NewMockTimeProvider(gomock.NewController(t))

	tp.EXPECT().CurrentTime(gomock.Any()).Return(now)
	tp.EXPECT().Since(gomock.Any(), now).Return(since)

	ch := latency.New(
		factory.NewClientFactory(),
		factory.NewIPApiFactory(),
		tp,
	)

	res, err := ch.Run(ctx, aggregate.NewTask(tk, px))
	require.NoError(t, err)
	require.Equal(t, since/2, res.LatencyResult.LatencyToProxy)
	require.Equal(t, since, res.LatencyResult.LatencyToTarget)
}

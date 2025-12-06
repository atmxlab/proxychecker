package checker_test

import (
	"context"
	"testing"
	stdtime "time"

	"github.com/atmxlab/proxychecker/internal/details/checker"
	"github.com/atmxlab/proxychecker/internal/details/checker/mocks"
	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/details/service/ipapi"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLatency(t *testing.T) {
	t.Parallel()

	var (
		ctx    = context.Background()
		tk     = entity.NewTaskBuilder().Build()
		px     = entity.NewProxyBuilder().Build()
		now    = stdtime.Now()
		since1 = 50 * stdtime.Millisecond
		since2 = 100 * stdtime.Millisecond
	)

	ctrl := gomock.NewController(t)

	cfg := client.Config{}
	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().Get(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, url string) ([]byte, error) {
			if hook := cfg.AfterDialHook(); hook != nil {
				hook()
			}

			return nil, nil
		})

	mockClientFactory := mocks.NewMockClientFactory(ctrl)
	mockClientFactory.EXPECT().Create(gomock.Any()).
		DoAndReturn(func(opts ...client.Option) client.Client {
			for _, opt := range opts {
				opt(&cfg)
			}

			return mockClient
		})
	mockClientFactory.EXPECT().Create(gomock.Any()).
		DoAndReturn(func(opts ...client.Option) client.Client {
			client.WithAfterDialHook(nil)(&cfg)
			return mockClient
		})

	mockIPApiClient := mocks.NewMockIPApi(ctrl)
	mockIPApiClient.EXPECT().Get(gomock.Any()).Return(ipapi.Output{}, nil)

	mockIPApiClientFactory := mocks.NewMockIPApiFactory(ctrl)
	mockIPApiClientFactory.EXPECT().Create(gomock.Any()).Return(mockIPApiClient)

	tp := mocks.NewMockTimeProvider(ctrl)

	tp.EXPECT().CurrentTime(gomock.Any()).Return(now)
	tp.EXPECT().Since(gomock.Any(), now).Return(since1)

	tp.EXPECT().CurrentTime(gomock.Any()).Return(now)
	tp.EXPECT().Since(gomock.Any(), now).Return(since2)

	ch := checker.NewLatencyChecker(
		mockClientFactory,
		mockIPApiClientFactory,
		tp,
	)

	res, err := ch.Run(ctx, aggregate.NewTask(tk, px))
	require.NoError(t, err)
	require.Equal(t, since1, res.LatencyResult.FromHostToProxyRoundTrip)
	require.Equal(t, since2, res.LatencyResult.FromHostToTargetRoundTrip)
}

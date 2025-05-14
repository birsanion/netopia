package handlers

import (
	"context"
	"time"

	"github.com/avast/retry-go/v4"
)

const (
	retryTaskDelay  = 10 * time.Second
	retryTaskJitter = 0.2
	retryMaxTries   = 5
	retryDelay      = time.Second
	retryMaxDelay   = 30 * time.Second
)

func RetryDo(ctx context.Context, fun retry.RetryableFunc, opts ...retry.Option) error {
	retryOpts := append([]retry.Option{
		retry.Context(ctx),
		retry.LastErrorOnly(true),
		retry.Attempts(retryMaxTries),
		retry.Delay(retryDelay),
		retry.MaxDelay(retryMaxDelay),
	}, opts...)
	return retry.Do(fun, retryOpts...)
}

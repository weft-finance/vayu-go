package client

import (
	"context"
	"time"
)

const ConnectionTimeoutAfterSeconds = 60

func GenerateContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ConnectionTimeoutAfterSeconds*time.Second)
}

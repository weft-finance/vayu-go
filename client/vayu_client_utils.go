package client

import (
	"context"
	"time"
)

const defaultTimeout = 60

var connectionTimeoutAfterSeconds = getConnectionTimeout()

func getConnectionTimeout() time.Duration {
	return defaultTimeout
}

func setConnectionTimeout(timeout time.Duration) {
	connectionTimeoutAfterSeconds = timeout
}

func GenerateContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), connectionTimeoutAfterSeconds*time.Second)
}

package client

import (
	"context"
	"time"
)

const connectionTimeout = 60 * time.Second

func GenerateContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), connectionTimeout)
}

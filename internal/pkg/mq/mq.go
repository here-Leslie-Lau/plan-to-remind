package mq

import (
	"context"
	"time"
)

type Producer interface {
	SendMsg(ctx context.Context, msg []byte) error
	DelayAtSendMsg(ctx context.Context, msg []byte, date *time.Time) error
	DelayAfterSendMsg(ctx context.Context, msg []byte, dur time.Duration) error
	CloseFunc()
}

type Consumer interface {
	Consume(ctx context.Context) ([]byte, error)
	CloseFunc()
}

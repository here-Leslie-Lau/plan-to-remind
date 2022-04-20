package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"plan-to-remind/internal/pkg/mq"
	"strings"
	"time"
)

type Consumer struct {
	consumers map[string]mq.Consumer
	log       *log.Helper
	handles   map[string]ConsumeHandle
}

type ConsumeHandle interface {
	Handle(ctx context.Context, result []byte) error
}

func (c *Consumer) Start(_ context.Context) error {
	// 消息监听
	nameList := []string{"plan_to_remind"}
	for _, name := range nameList {
		go c.listenConsumer(name)
	}
	return nil
}

func (c *Consumer) Stop(_ context.Context) error {
	// close func
	for _, consumer := range c.consumers {
		consumer.CloseFunc()
	}
	return nil
}

func (c *Consumer) listenConsumer(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for {
		result, err := c.consumers[name].Consume(context.Background())
		if err != nil {
			if strings.Contains(err.Error(), "ConsumerClosed") {
				c.log.Warnw("listenConsumer consume close", "name:", name, "err:", err)
				return
			}
			c.log.Errorw("listenConsumer consume fail", "err:", err)
			continue
		}
		if err := c.handles[name].Handle(ctx, result); err != nil {
			c.log.Errorw("listenConsumer Handle fail", "name:", name, "result:", string(result), "err:", err)
			continue
		}
	}
}

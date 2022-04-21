package server

import (
	"context"
	"plan-to-remind/internal/pkg/log"
	"plan-to-remind/internal/pkg/mq"
	"strings"
)

type Consumer struct {
	consumers map[string]mq.Consumer
	log       *log.Helper
	handles   map[string]Handler
}

type Handler func(ctx context.Context, result []byte) error

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
	for {
		consumer, ok := c.consumers[name]
		if !ok {
			c.log.Error("listenConsumer consumers is nil", "name", name)
			continue
		}
		result, err := consumer.Consume(context.Background())
		if err != nil {
			if strings.Contains(err.Error(), "ConsumerClosed") {
				c.log.Warn("listenConsumer consume close", "name:", name, "err:", err)
				return
			}
			c.log.Error("listenConsumer consume fail", "err:", err)
			continue
		}
		handler, ok := c.handles[name]
		if !ok {
			c.log.Error("listenConsumer handle is nil", "name", name)
			continue
		}
		if err := handler(context.Background(), result); err != nil {
			c.log.Error("listenConsumer Handle fail", "name:", name, "result:", string(result), "err:", err)
			continue
		}
	}
}

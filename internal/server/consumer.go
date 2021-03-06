package server

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	klog "github.com/go-kratos/kratos/v2/log"
	"plan-to-remind/internal/conf"
	"plan-to-remind/internal/pkg/log"
	"plan-to-remind/internal/pkg/mq"
	consumerpkg "plan-to-remind/internal/pkg/mq/consumer"
	"plan-to-remind/internal/service"
	"strings"
)

const (
	handleNamePlanToRemind = "plan_to_remind"
)

var handleNames = []string{handleNamePlanToRemind}

type ConsumerServer struct {
	consumer  *service.ConsumerService
	consumers map[string]mq.Consumer
	log       *log.Helper
	handles   map[string]Handler
}

func NewConsumerServer(consumerService *service.ConsumerService, data *conf.Data, logger klog.Logger) *ConsumerServer {
	cli, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: data.Pulsar.Url,
	})
	if err != nil {
		panic(err)
	}

	consumers := make(map[string]mq.Consumer)
	for _, name := range handleNames {
		consumer, err := consumerpkg.NewPulsarConsumer(data.Pulsar.Topic, name, cli)
		if err != nil {
			panic(err)
		}
		consumers[name] = consumer
	}

	// a little ugly...
	handles := make(map[string]Handler)
	handles[handleNamePlanToRemind] = consumerService.PlanToRemind
	return &ConsumerServer{log: log.NewHelper(logger), handles: handles, consumer: consumerService, consumers: consumers}
}

type Handler func(ctx context.Context, result []byte) error

func (c *ConsumerServer) Start(_ context.Context) error {
	// 消息监听
	for _, name := range handleNames {
		go c.listenConsumer(name)
	}
	c.log.Debug("ConsumerServer listenConsumer success", "names", handleNames)
	return nil
}

func (c *ConsumerServer) Stop(_ context.Context) error {
	// close func
	for _, consumer := range c.consumers {
		consumer.CloseFunc()
	}
	return nil
}

func (c *ConsumerServer) listenConsumer(name string) {
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
			continue
		}
	}
}

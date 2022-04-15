package biz

import (
	"context"
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	"plan-to-remind/internal/conf"
	"plan-to-remind/internal/data/model"
	"plan-to-remind/internal/pkg/mq"
	producer2 "plan-to-remind/internal/pkg/mq/producer"
	"time"
)

type TimerUsecase struct {
	parser   *cronParser
	producer mq.Producer
}

func NewTimerUsecase(data *conf.Data) *TimerUsecase {
	cli, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: data.Pulsar.Url,
	})
	if err != nil {
		panic(err)
	}
	producer, err := producer2.NewPulsarProducer(data.Pulsar.Topic, cli)
	if err != nil {
		panic(err)
	}
	return &TimerUsecase{parser: &cronParser{}, producer: producer}
}

// UserPlanPush 用户任务推送
func (t *TimerUsecase) UserPlanPush(ctx context.Context) error {
	// find all active plans
	plan := NewDefaultPlan(0)
	list, err := plan.List(ctx, &PlanFilter{State: model.PlanStateOn})
	if err != nil {
		return err
	}
	// push plan to delay queue
	for _, p := range list {
		bytes, err := json.Marshal(p)
		if err != nil {
			// todo log
			continue
		}
		dur := t.parser.Parser(p.CronDesc)
		if err := t.producer.DelayAfterSendMsg(ctx, bytes, dur); err != nil {
			// todo log
			continue
		}
	}
	return nil
}

type cronParser struct{}

func (c *cronParser) Parser(expression string) time.Duration {
	return 0
}

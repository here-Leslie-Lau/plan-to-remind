package biz

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/go-kratos/kratos/v2/log"
	"plan-to-remind/internal/conf"
	"plan-to-remind/internal/data/model"
	"plan-to-remind/internal/pkg/json"
	"plan-to-remind/internal/pkg/mq"
	producer2 "plan-to-remind/internal/pkg/mq/producer"
	"strconv"
	"strings"
	"time"
)

type TimerUsecase struct {
	parser   *cronParser
	producer mq.Producer
	log      *log.Helper
	json     *json.Parser
}

func NewTimerUsecase(data *conf.Data, logger log.Logger, json *json.Parser) (*TimerUsecase, func()) {
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
	return &TimerUsecase{parser: &cronParser{}, producer: producer, json: json, log: log.NewHelper(logger)}, producer.CloseFunc
}

// UserPlanPush 用户任务推送
func (uc *TimerUsecase) UserPlanPush(ctx context.Context) error {
	// find all active plans
	plan := NewDefaultPlan(0)
	list, err := plan.List(ctx, &PlanFilter{State: model.PlanStateOn})
	if err != nil {
		return err
	}
	// push plan to delay queue
	for _, p := range list {
		bytes, err := uc.json.Marshal(p)
		if err != nil {
			uc.log.Errorw("UserPlanPush json marshal fail, bytes:%s, err:%v", string(bytes), err)
			continue
		}
		dur := uc.parser.Parser(p.CronDesc)
		if dur == 0 {
			uc.log.Debugw("UserPlanPush Parser cron_expression continue", "duration:", dur)
			continue
		}
		if err := uc.producer.DelayAfterSendMsg(ctx, bytes, dur); err != nil {
			uc.log.Errorw("UserPlanPush DelayAfterSendMsg fail, bytes:%s, dur:%v, err:%v", string(bytes), dur, err)
			continue
		}
	}
	return nil
}

type cronParser struct{}

// Parser 简陋版时间解析器实现,仅仅扫描当日的,todo 后续迭代更多功能
func (c *cronParser) Parser(expression string) time.Duration {
	now := time.Now()
	list := strings.Split(expression, " ")
	if len(list) != 5 {
		return 0
	}
	// 仅仅扫描当日的
	if list[2] != "*" && list[2] != strconv.FormatInt(int64(now.Day()), 10) {
		return 0
	}

	hour, err := strconv.ParseInt(list[1], 10, 64)
	if err != nil {
		return 0
	}
	min, err := strconv.ParseInt(list[0], 10, 64)
	if err != nil {
		return 0
	}
	dur := time.Date(now.Year(), now.Month(), now.Day(), int(hour), int(min), 0, 0, now.Location())
	return dur.Sub(now)
}

func (uc *TimerUsecase) PlanToInvalid(ctx context.Context) error {
	plan := NewDefaultPlan(0)
	list, err := plan.List(ctx, &PlanFilter{
		State:       model.PlanStateOn,
		DeadTimeEnd: time.Now().Unix(),
	})
	if err != nil {
		return err
	}
	for _, p := range list {
		p.repo = RepoPlan
		p.State = model.PlanStateOff
		if err := p.Update(ctx); err != nil {
			uc.log.Errorw("PlanToInvalid update plan error", "plan_id:", p.ID, "err:", err)
			continue
		}
	}
	return nil
}

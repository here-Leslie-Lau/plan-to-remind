package biz

import (
	"context"
	"plan-to-remind/internal/data/model"
	"time"
)

type TimerUsecase struct {
	parser *cronParser
}

func NewTimerUsecase() *TimerUsecase {
	return &TimerUsecase{parser: &cronParser{}}
}

// UserPlanPush 用户任务推送
func (t *TimerUsecase) UserPlanPush(ctx context.Context) error {
	// find all active plans
	plan := NewDefaultPlan(0)
	_, err := plan.List(ctx, &PlanFilter{State: model.PlanStateOn})
	if err != nil {
		return err
	}
	// push plan to delay queue
	return nil
}

type cronParser struct{}

func (c *cronParser) Parser(expression string) time.Duration {
	return 0
}

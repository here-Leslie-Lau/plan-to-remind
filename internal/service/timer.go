package service

import (
	"context"
	klog "github.com/go-kratos/kratos/v2/log"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/pkg/log"
	"time"
)

type TimerService struct {
	log *log.Helper
	uc  *biz.TimerUsecase
}

func NewTimerService(logger klog.Logger, uc *biz.TimerUsecase) *TimerService {
	return &TimerService{log: log.NewHelper(logger), uc: uc}
}

func (t *TimerService) UserPlanPush() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := t.uc.UserPlanPush(ctx); err != nil {
		t.log.Error("UserPlanPush usecase UserPlanPush error", "err:", err)
	}
}

// PlanToInvalid 将计划失效
func (t *TimerService) PlanToInvalid() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := t.uc.PlanToInvalid(ctx); err != nil {
		t.log.Error("PlanToInvalid usecase PlanToInvalid error", "time:", time.Now(), "err:", err)
	}
}

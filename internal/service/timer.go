package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"plan-to-remind/internal/biz"
	"time"
)

type TimerService struct {
	log *log.Helper
	uc  *biz.TimerUsecase
}

func NewTimerService(logger log.Logger, uc *biz.TimerUsecase) *TimerService {
	return &TimerService{log: log.NewHelper(logger), uc: uc}
}

func (t *TimerService) UserPlanPush() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := t.uc.UserPlanPush(ctx); err != nil {
		t.log.Errorw("UserPlanPush usecase UserPlanPush error", "err:", err)
	}
}

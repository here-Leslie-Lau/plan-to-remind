package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"plan-to-remind/internal/biz"
)

type TimerService struct {
	log *log.Helper
	uc  *biz.TimerUsecase
}

func NewTimerService(logger log.Logger, uc *biz.TimerUsecase) *TimerService {
	return &TimerService{log: log.NewHelper(logger), uc: uc}
}

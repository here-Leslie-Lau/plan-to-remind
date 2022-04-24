package service

import (
	"context"
	klog "github.com/go-kratos/kratos/v2/log"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/pkg/log"
)

type ConsumerService struct {
	uc  *biz.ConsumerUsecase
	log *log.Helper
}

func NewConsumerService(uc *biz.ConsumerUsecase, logger klog.Logger) *ConsumerService {
	return &ConsumerService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ConsumerService) PlanToRemind(ctx context.Context, result []byte) error {
	if err := s.uc.PlanToRemind(ctx, result); err != nil {
		s.log.Error("PlanToRemind consumer usecase PlanToRemind fail", "result", string(result), "err", err)
		return err
	}
	return nil
}

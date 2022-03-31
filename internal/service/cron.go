package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"plan-to-remind/internal/biz"

	v1 "plan-to-remind/api/plantoremind/v1/cron"
)

type CronService struct {
	v1.UnimplementedCronServer

	uc  *biz.CronSpecUsecase
	log *log.Helper
}

func NewCronService(uc *biz.CronSpecUsecase, logger log.Logger) *CronService {
	return &CronService{uc: uc, log: log.NewHelper(logger)}
}

func (s *CronService) CreateCron(ctx context.Context, req *v1.CreateCronRequest) (*emptypb.Empty, error) {
	err := s.uc.CreateCronSpec(ctx, &biz.CronSpec{
		Desc:       req.Desc,
		Expression: req.Expression,
	})
	if err != nil {
		s.log.Errorw("CreateCron CreateCronSpec error", "req:", req, "err:", err)
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *CronService) UpdateCron(ctx context.Context, req *v1.UpdateCronRequest) (*emptypb.Empty, error) {
	err := s.uc.UpdateCronSpec(ctx, &biz.CronSpec{
		ID:         req.Id,
		Desc:       req.Desc,
		Expression: req.Expression,
	})
	if err != nil {
		s.log.Errorw("UpdateCron UpdateCronSpec error", "req:", req, "err:", err)
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *CronService) DeleteCron(ctx context.Context, req *v1.DeleteCronRequest) (*v1.DeleteCronReply, error) {
	return &v1.DeleteCronReply{}, nil
}

func (s *CronService) GetCron(ctx context.Context, req *v1.GetCronRequest) (*v1.GetCronReply, error) {
	cronSpec, err := s.uc.GetCronSpec(ctx, req.Id)
	if err != nil {
		s.log.Warnw("GetCron GetCronSpec error", "id:", req.Id, "err:", err)
		return nil, err
	}
	data := &v1.CronData{
		Id:         cronSpec.ID,
		Desc:       cronSpec.Desc,
		Expression: cronSpec.Expression,
	}
	return &v1.GetCronReply{Data: data}, nil
}

func (s *CronService) ListCron(ctx context.Context, req *v1.ListCronRequest) (*v1.ListCronReply, error) {
	return &v1.ListCronReply{}, nil
}

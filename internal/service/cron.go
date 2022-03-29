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

func (s *CronService) UpdateCron(ctx context.Context, req *v1.UpdateCronRequest) (*v1.UpdateCronReply, error) {
	return &v1.UpdateCronReply{}, nil
}

func (s *CronService) DeleteCron(ctx context.Context, req *v1.DeleteCronRequest) (*v1.DeleteCronReply, error) {
	return &v1.DeleteCronReply{}, nil
}

func (s *CronService) GetCron(ctx context.Context, req *v1.GetCronRequest) (*v1.GetCronReply, error) {
	return &v1.GetCronReply{}, nil
}

func (s *CronService) ListCron(ctx context.Context, req *v1.ListCronRequest) (*v1.ListCronReply, error) {
	return &v1.ListCronReply{}, nil
}

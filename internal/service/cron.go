package service

import (
	"context"
	klog "github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/pkg/log"

	v1 "plan-to-remind/api/plantoremind/v1/cron"
)

type CronService struct {
	v1.UnimplementedCronServer

	uc  *biz.CronSpecUsecase
	log *log.Helper
}

func NewCronService(uc *biz.CronSpecUsecase, logger klog.Logger) *CronService {
	return &CronService{uc: uc, log: log.NewHelper(logger)}
}

func (s *CronService) CreateCron(ctx context.Context, req *v1.CreateCronRequest) (*emptypb.Empty, error) {
	cron := biz.NewCronSpec(0, req.Desc, req.Expression)
	err := s.uc.CreateCronSpec(ctx, cron)
	if err != nil {
		s.log.Error("CreateCron CreateCronSpec error", "req:", req, "err:", err)
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *CronService) UpdateCron(ctx context.Context, req *v1.UpdateCronRequest) (*emptypb.Empty, error) {
	cron := biz.NewCronSpec(req.Id, req.Desc, req.Expression)
	err := s.uc.UpdateCronSpec(ctx, cron)
	if err != nil {
		s.log.Error("UpdateCron UpdateCronSpec error", "req:", req, "err:", err)
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *CronService) DeleteCron(ctx context.Context, req *v1.DeleteCronRequest) (*emptypb.Empty, error) {
	if err := s.uc.DeleteCronSpec(ctx, req.Id); err != nil {
		s.log.Error("DeleteCron DeleteCronSpec error", "id:", req.Id, "err;", err)
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *CronService) GetCron(ctx context.Context, req *v1.GetCronRequest) (*v1.GetCronReply, error) {
	cronSpec, err := s.uc.GetCronSpec(ctx, req.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.log.Warn("GetCron GetCronSpec error", "id:", req.Id, "err:", err)
		return nil, err
	}
	if cronSpec == nil {
		return new(v1.GetCronReply), nil
	}
	data := &v1.CronData{
		Id:         cronSpec.ID,
		Desc:       cronSpec.Desc,
		Expression: cronSpec.Expression,
	}
	return &v1.GetCronReply{Data: data}, nil
}

func (s *CronService) ListCron(ctx context.Context, req *v1.ListCronRequest) (*v1.ListCronReply, error) {
	f := &biz.CronSpecFilter{
		OffSet:  req.Offset,
		Limit:   req.Limit,
		OrderBy: req.OrderBy,
	}
	cronSpec, err := s.uc.ListCronSpec(ctx, f)
	if err != nil {
		s.log.Warn("ListCron ListCronSpec error", "req:", req, "err:", err)
		return nil, err
	}

	var list []*v1.CronData
	for _, spec := range cronSpec {
		list = append(list, &v1.CronData{
			Id:         spec.ID,
			Desc:       spec.Desc,
			Expression: spec.Expression,
		})
	}
	return &v1.ListCronReply{List: list}, nil
}

package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/pkg/format"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	v1 "plan-to-remind/api/plantoremind/v1/plan"
)

type PlanService struct {
	v1.UnimplementedPlanServer

	uc  *biz.PlanUsecase
	log *log.Helper
}

func NewPlanService(uc *biz.PlanUsecase, logger log.Logger) *PlanService {
	return &PlanService{uc: uc, log: log.NewHelper(logger)}
}

func (s *PlanService) CreatePlan(ctx context.Context, req *v1.CreatePlanRequest) (*emptypb.Empty, error) {
	deadTime, err := time.Parse(format.DateSecondLayout, req.DeadTime)
	if err != nil {
		s.log.Warnw("CreatePlan time parse error", "dead_time:", req.DeadTime, "err:", err)
		return nil, err
	}
	err = s.uc.CreatePlan(ctx, &biz.Plan{
		State:    uint8(req.State),
		Level:    uint8(req.Level),
		CronId:   req.CronId,
		DeadTime: deadTime.Unix(),
		Name:     req.Name,
	})
	if err != nil {
		s.log.Errorw("CreatePlan biz.CreatePlan err", "req:", req, "err:", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *PlanService) UpdatePlan(ctx context.Context, req *v1.UpdatePlanRequest) (*v1.UpdatePlanReply, error) {
	return &v1.UpdatePlanReply{}, nil
}
func (s *PlanService) DeletePlan(ctx context.Context, req *v1.DeletePlanRequest) (*v1.DeletePlanReply, error) {
	return &v1.DeletePlanReply{}, nil
}
func (s *PlanService) GetPlan(ctx context.Context, req *v1.GetPlanRequest) (*v1.GetPlanReply, error) {
	plan, err := s.uc.GetPlanByID(ctx, req.Id)
	if err != nil {
		s.log.Warnw("GetPlan GetPlanByID error", "plan_id:", req.Id, "err:", err)
		return nil, err
	}
	return &v1.GetPlanReply{Plan: &v1.PlanData{
		Id:       plan.ID,
		State:    uint32(plan.State),
		Level:    uint32(plan.Level),
		CronId:   plan.CronId,
		DeadTime: plan.DeadTime,
		Name:     plan.Name,
		CronDesc: plan.CronDesc,
	}}, nil
}
func (s *PlanService) ListPlan(ctx context.Context, req *v1.ListPlanRequest) (*v1.ListPlanReply, error) {
	return &v1.ListPlanReply{}, nil
}

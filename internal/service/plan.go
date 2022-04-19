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
	deadTime, err := time.Parse(format.DateLayout, req.DeadTime)
	if err != nil {
		s.log.Warnw("CreatePlan time parse error", "dead_time:", req.DeadTime, "err:", err)
		return nil, err
	}
	plan := biz.NewPlan(0, uint8(req.State), uint8(req.Level), req.CronId, deadTime.Unix(), req.Name, "")
	err = s.uc.CreatePlan(ctx, plan)
	if err != nil {
		s.log.Errorw("CreatePlan biz.CreatePlan err", "req:", req, "err:", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *PlanService) UpdatePlan(ctx context.Context, req *v1.UpdatePlanRequest) (*emptypb.Empty, error) {
	plan := biz.NewPlan(req.Id, uint8(req.State), uint8(req.Level), req.CronId, 0, req.Name, "")

	if req.DeadTime != "" {
		deadTime, err := time.Parse(req.DeadTime, format.DateLayout)
		if err != nil {
			s.log.Warnw("UpdatePlan time parse dead_time error", "dead_time:", req.DeadTime, "err:", err)
			return nil, err
		}
		plan.DeadTime = deadTime.Unix()
	}

	if err := s.uc.UpdatePlan(ctx, plan); err != nil {
		s.log.Errorw("UpdatePlan biz UpdatePlan error", "req:", req, "err:", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *PlanService) DeletePlan(ctx context.Context, req *v1.DeletePlanRequest) (*emptypb.Empty, error) {
	if err := s.uc.DeletePlan(ctx, req.Id); err != nil {
		s.log.Errorw("DeletePlan biz.DeletePlan error", "plan_id:", req.Id, "err:", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
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
	f := &biz.PlanFilter{
		Offset:  req.Offset,
		Limit:   req.Limit,
		OrderBy: req.OrderBy,
	}
	if req.DeadTimeBegin != "" {
		begin, err := time.Parse(format.DateLayout, req.DeadTimeBegin)
		if err != nil {
			s.log.Warnw("ListPlan time parse dead_time_begin error", "begin:", req.DeadTimeBegin, "err:", err)
			return nil, err
		}
		f.DeadTimeBegin = begin.Unix()
	}
	if req.DeadTimeEnd != "" {
		end, err := time.Parse(format.DateLayout, req.DeadTimeEnd)
		if err != nil {
			s.log.Warnw("ListPlan time parse dead_time_end error", "begin:", req.DeadTimeBegin, "err:", err)
			return nil, err
		}
		f.DeadTimeEnd = end.Unix()
	}
	plans, err := s.uc.ListPlanByFilter(ctx, f)
	if err != nil {
		s.log.Warnw("ListPlan ListPlanByFilter error", "req:", req, "err:", err)
		return nil, err
	}

	var list []*v1.PlanData
	for _, plan := range plans {
		list = append(list, &v1.PlanData{
			Id:       plan.ID,
			State:    uint32(plan.State),
			Level:    uint32(plan.Level),
			CronId:   plan.CronId,
			DeadTime: plan.DeadTime,
			Name:     plan.Name,
			CronDesc: plan.CronDesc,
		})
	}
	return &v1.ListPlanReply{List: list}, nil
}

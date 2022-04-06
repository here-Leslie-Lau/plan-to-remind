package biz

import (
	"context"
	"time"
)

// Plan 计划表
type Plan struct {
	ID uint64 `json:"id"`
	// 计划开关状态
	State uint8 `json:"state"`
	// 计划等级
	Level uint8 `json:"level"`
	// 计划时间表ID
	CronId uint64 `json:"cron_id"`
	// 计划失效时间
	DeadTime int64 `json:"dead_time"`
	// 计划名称
	Name string `json:"name"`
	// 计划时间表描述
	CronDesc string `json:"cron_desc"`
}

type PlanFilter struct {
	DeadTimeBegin int64
	DeadTimeEnd   int64
	// 分页
	Offset, Limit         int64
	TotalCount, TotalPage int64
	OrderBy               string
}

type PlanRepo interface {
	CreatePlan(ctx context.Context, plan *Plan) error
	GetPlan(ctx context.Context, id uint64) (*Plan, error)
	UpdatePlan(ctx context.Context, id uint64, param map[string]interface{}) error
	ListPlanByFilter(ctx context.Context, f *PlanFilter) ([]*Plan, error)
}

type PlanUsecase struct {
	repo PlanRepo
}

func NewPlanUsecase(repo PlanRepo) *PlanUsecase {
	return &PlanUsecase{repo: repo}
}

func (uc *PlanUsecase) GetPlanByID(ctx context.Context, id uint64) (*Plan, error) {
	return uc.repo.GetPlan(ctx, id)
}

func (uc *PlanUsecase) CreatePlan(ctx context.Context, plan *Plan) error {
	return uc.repo.CreatePlan(ctx, plan)
}

func (uc *PlanUsecase) UpdatePlan(ctx context.Context, plan *Plan) error {
	param := make(map[string]interface{})
	if plan.State > 0 {
		param["state"] = plan.State
	}
	if plan.Level > 0 {
		param["level"] = plan.Level
	}
	if plan.CronId > 0 {
		param["cron_id"] = plan.CronId
	}
	if plan.DeadTime > 0 {
		param["dead_time"] = plan.DeadTime
	}
	if plan.Name != "" {
		param["name"] = plan.Name
	}
	return uc.repo.UpdatePlan(ctx, plan.ID, param)
}

func (uc *PlanUsecase) DeletePlan(ctx context.Context, id uint64) error {
	param := map[string]interface{}{
		"deleted_at": time.Now(),
	}
	return uc.repo.UpdatePlan(ctx, id, param)
}

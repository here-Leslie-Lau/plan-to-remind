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
	repo     PlanRepo
}

func NewPlan(id uint64, state uint8, level uint8, cronId uint64, deadTime int64, name string, cronDesc string) *Plan {
	return &Plan{ID: id, State: state, Level: level, CronId: cronId, DeadTime: deadTime, Name: name, CronDesc: cronDesc, repo: RepoPlan}
}

func NewDefaultPlan(id uint64) *Plan {
	return &Plan{
		ID:   id,
		repo: RepoPlan,
	}
}

type PlanFilter struct {
	DeadTimeBegin int64
	DeadTimeEnd   int64
	// 分页
	Offset, Limit int64
	OrderBy       string
}

type PlanRepo interface {
	CreatePlan(ctx context.Context, plan *Plan) error
	GetPlan(ctx context.Context, id uint64) (*Plan, error)
	UpdatePlan(ctx context.Context, id uint64, param map[string]interface{}) error
	ListPlanByFilter(ctx context.Context, f *PlanFilter) ([]*Plan, error)
}

var RepoPlan PlanRepo

func (p *Plan) Get(ctx context.Context) (*Plan, error) {
	return p.repo.GetPlan(ctx, p.ID)
}

func (p *Plan) Create(ctx context.Context) error {
	return p.repo.CreatePlan(ctx, p)
}

func (p *Plan) Update(ctx context.Context) error {
	param := make(map[string]interface{})
	if p.State > 0 {
		param["state"] = p.State
	}
	if p.Level > 0 {
		param["level"] = p.Level
	}
	if p.CronId > 0 {
		param["cron_id"] = p.CronId
	}
	if p.DeadTime > 0 {
		param["dead_time"] = p.DeadTime
	}
	if p.Name != "" {
		param["name"] = p.Name
	}
	return p.repo.UpdatePlan(ctx, p.ID, param)
}

func (p *Plan) Delete(ctx context.Context) error {
	param := map[string]interface{}{
		"deleted_at": time.Now(),
	}
	return p.repo.UpdatePlan(ctx, p.ID, param)
}

func (p *Plan) List(ctx context.Context, f *PlanFilter) ([]*Plan, error) {
	return p.repo.ListPlanByFilter(ctx, f)
}

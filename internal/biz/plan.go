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
	// cron expression
	Expression string `json:"expression"`
	repo       PlanRepo
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
	State         uint8
	DeadTimeBegin int64
	DeadTimeEnd   int64
	// 分页
	Offset, Limit int64
	OrderBy       string
}

type PlanCompletion struct {
	BeginAt int64 `json:"begin_at"`
	EndAt   int64 `json:"end_at"`
	// value object
	PlanId        uint64 `json:"plan_id"`
	TotalNums     int    `json:"total_nums"`
	CompletedNums int    `json:"completed_nums"`
}

func NewDefaultPlanCompletion(planId uint64) *PlanCompletion {
	now := time.Now()
	// 默认一周
	return &PlanCompletion{
		BeginAt:   now.Unix(),
		EndAt:     now.AddDate(0, 0, 7).Unix(),
		PlanId:    planId,
		TotalNums: 7,
	}
}

type PlanRepo interface {
	SavePlan(ctx context.Context, plan *Plan) error
	GetPlan(ctx context.Context, id uint64) (*Plan, error)
	DeletePlan(ctx context.Context, id uint64) error
	ListPlanByFilter(ctx context.Context, f *PlanFilter) ([]*Plan, error)
	SavePlanCompletion(ctx context.Context, completion *PlanCompletion) error
	GetLatestPlanCompletion(ctx context.Context, planId uint64) (*PlanCompletion, error)
}

var RepoPlan PlanRepo

func (p *Plan) Get(ctx context.Context) (*Plan, error) {
	return p.repo.GetPlan(ctx, p.ID)
}

func (p *Plan) Create(ctx context.Context) error {
	if err := p.repo.SavePlan(ctx, p); err != nil {
		return err
	}
	return p.repo.SavePlanCompletion(ctx, NewDefaultPlanCompletion(p.ID))
}

func (p *Plan) Update(ctx context.Context) error {
	return p.repo.SavePlan(ctx, p)
}

func (p *Plan) Delete(ctx context.Context) error {
	return p.repo.DeletePlan(ctx, p.ID)
}

func (p *Plan) List(ctx context.Context, f *PlanFilter) ([]*Plan, error) {
	return p.repo.ListPlanByFilter(ctx, f)
}

func (p *Plan) Complete(ctx context.Context) error {
	completion, err := p.repo.GetLatestPlanCompletion(ctx, p.ID)
	if err != nil {
		return err
	}
	completion.CompletedNums += 1
	return p.repo.SavePlanCompletion(ctx, completion)
}

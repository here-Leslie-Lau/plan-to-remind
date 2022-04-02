package biz

import "context"

// Plan 计划表
type Plan struct {
	ID uint64 `json:"id"`
	// 计划开关状态
	State uint8 `json:"state"`
	// 计划等级
	Level uint8 `json:"level"`
	// 计划时间表ID
	DescId uint64 `json:"desc_id"`
	// 计划失效时间
	DeadTime int64 `json:"dead_time"`
	// 计划名称
	Name string `json:"name"`
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

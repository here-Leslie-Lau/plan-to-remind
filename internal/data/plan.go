package data

import (
	"context"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/data/model"
)

// PlanRepo biz.PlanRepo实现，计划表仓库
type PlanRepo struct {
	data *Data
}

func (p *PlanRepo) CreatePlan(ctx context.Context, plan *biz.Plan) error {
	panic("implement me")
}

func (p *PlanRepo) GetPlan(ctx context.Context, id uint64) (*biz.Plan, error) {
	plan := &model.Plan{}
	if err := p.data.WithCtx(ctx).Model(&model.Plan{}).Where("id = ?", id).First(plan).Error; err != nil {
		return nil, err
	}
	return &biz.Plan{
		ID:       uint64(plan.ID),
		State:    plan.State,
		Level:    plan.Level,
		DescId:   plan.DescId,
		DeadTime: plan.DeadTime,
		Name:     plan.Name,
	}, nil
}

func (p *PlanRepo) UpdatePlan(ctx context.Context, id uint64, param map[string]interface{}) error {
	panic("implement me")
}

func (p *PlanRepo) ListPlanByFilter(ctx context.Context, f *biz.PlanFilter) ([]*biz.Plan, error) {
	panic("implement me")
}

func NewPlanRepo(data *Data) *PlanRepo {
	return &PlanRepo{data: data}
}

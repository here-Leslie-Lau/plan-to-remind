package data

import (
	"context"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/data/model"
)

var _ biz.PlanRepo = (*PlanRepo)(nil)

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
	result := &biz.Plan{
		ID:       uint64(plan.ID),
		State:    plan.State,
		Level:    plan.Level,
		CronId:   plan.DescId,
		DeadTime: plan.DeadTime,
		Name:     plan.Name,
	}
	// query cron desc
	var desc string
	err := p.data.WithCtx(ctx).Model(&model.CronSpec{}).Select("desc").Where("id = ?", id).Scan(&desc).Error
	if err != nil {
		return nil, err
	}
	result.CronDesc = desc
	return result, nil
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

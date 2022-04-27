package biz

import (
	"context"
)

type PlanUsecase struct{}

func NewPlanUsecase() *PlanUsecase {
	return &PlanUsecase{}
}

func (uc *PlanUsecase) GetPlanByID(ctx context.Context, id uint64) (*Plan, error) {
	plan := NewDefaultPlan(id)
	plan, err := plan.Get(ctx)
	if err != nil {
		return nil, err
	}
	// get cron by cron_id
	cron := NewDefaultCron(plan.CronId)
	cron, err = cron.Get(ctx)
	if err != nil {
		return nil, err
	}
	plan.CronDesc = cron.Desc
	plan.Expression = cron.Expression
	return plan, nil
}

func (uc *PlanUsecase) CreatePlan(ctx context.Context, plan *Plan) error {
	return plan.Create(ctx)
}

func (uc *PlanUsecase) UpdatePlan(ctx context.Context, plan *Plan) error {
	return plan.Update(ctx)
}

func (uc *PlanUsecase) DeletePlan(ctx context.Context, id uint64) error {
	plan := NewDefaultPlan(id)
	return plan.Delete(ctx)
}

func (uc *PlanUsecase) ListPlanByFilter(ctx context.Context, f *PlanFilter) ([]*Plan, error) {
	plan := NewDefaultPlan(0)
	return plan.List(ctx, f)
}

func (uc *PlanUsecase) CompletePlan(ctx context.Context, id uint64) error {
	plan := NewDefaultPlan(id)
	return plan.Complete(ctx)
}

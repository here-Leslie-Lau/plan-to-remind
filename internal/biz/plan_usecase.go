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
	return plan.Get(ctx)
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

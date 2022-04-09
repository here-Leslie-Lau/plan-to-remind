package biz

import (
	"context"
)

type CronSpecUsecase struct {
	cron *CronSpec
}

func NewCronSpecUsecase() *CronSpecUsecase {
	return &CronSpecUsecase{}
}

func (uc *CronSpecUsecase) CreateCronSpec(ctx context.Context, cron *CronSpec) error {
	return cron.Create(ctx)
}

func (uc *CronSpecUsecase) UpdateCronSpec(ctx context.Context, cron *CronSpec) error {
	return cron.Update(ctx)
}

func (uc *CronSpecUsecase) GetCronSpec(ctx context.Context, id uint64) (*CronSpec, error) {
	cron := NewDefaultCron(id)
	return cron.Get(ctx)
}

func (uc *CronSpecUsecase) DeleteCronSpec(ctx context.Context, id uint64) error {
	cron := NewDefaultCron(id)
	return cron.Delete(ctx)
}

func (uc *CronSpecUsecase) ListCronSpec(ctx context.Context, f *CronSpecFilter) ([]*CronSpec, error) {
	cron := NewDefaultCron(0)
	return cron.ListCronSpec(ctx, f)
}

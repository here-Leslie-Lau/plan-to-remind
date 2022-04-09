package biz

import (
	"context"
)

type CronSpecUsecase struct {
	repo CronSpecRepo
	cron *CronSpec
}

func NewCronSpecUsecase(repo CronSpecRepo) *CronSpecUsecase {
	return &CronSpecUsecase{repo: repo}
}

func (uc *CronSpecUsecase) CreateCronSpec(ctx context.Context, cron *CronSpec) error {
	return cron.Create(ctx)
}

func (uc *CronSpecUsecase) UpdateCronSpec(ctx context.Context, cron *CronSpec) error {
	return cron.Update(ctx)
}

func (uc *CronSpecUsecase) GetCronSpec(ctx context.Context, id uint64) (*CronSpec, error) {
	cron := NewCronSpec(id, "", "")
	return cron.Get(ctx)
}

func (uc *CronSpecUsecase) DeleteCronSpec(ctx context.Context, id uint64) error {
	cron := NewCronSpec(id, "", "")
	return cron.Delete(ctx)
}

func (uc *CronSpecUsecase) ListCronSpec(ctx context.Context, f *CronSpecFilter) ([]*CronSpec, error) {
	return uc.repo.ListCronSpec(ctx, f)
}

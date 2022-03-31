package biz

import (
	"context"
	"time"
)

type CronSpec struct {
	ID         uint64 `json:"id"`
	Desc       string `json:"desc"`
	Expression string `json:"expression"`
}

type CronSpecFilter struct {
	OffSet     int64  `json:"off_set"`
	Limit      int64  `json:"limit"`
	TotalCount int64  `json:"total_count"`
	TotalPage  int64  `json:"total_page"`
	OrderBy    string `json:"order_by"`
}

type CronSpecRepo interface {
	SaveCronSpec(ctx context.Context, cron *CronSpec) error
	GetCronSpec(ctx context.Context, id uint64) (*CronSpec, error)
	ListCronSpec(ctx context.Context, f *CronSpecFilter) ([]*CronSpec, error)
	UpdateCronSpec(ctx context.Context, id uint64, params map[string]interface{}) error
}

type CronSpecUsecase struct {
	repo CronSpecRepo
}

func NewCronSpecUsecase(repo CronSpecRepo) *CronSpecUsecase {
	return &CronSpecUsecase{repo: repo}
}

func (uc *CronSpecUsecase) CreateCronSpec(ctx context.Context, cron *CronSpec) error {
	return uc.repo.SaveCronSpec(ctx, cron)
}

func (uc *CronSpecUsecase) UpdateCronSpec(ctx context.Context, cron *CronSpec) error {
	param := make(map[string]interface{}, 3)
	if cron.Desc != "" {
		param["desc"] = cron.Desc
	}
	if cron.Expression != "" {
		param["expression"] = cron.Expression
	}
	return uc.repo.UpdateCronSpec(ctx, cron.ID, param)
}

func (uc *CronSpecUsecase) GetCronSpec(ctx context.Context, id uint64) (*CronSpec, error) {
	return uc.repo.GetCronSpec(ctx, id)
}

func (uc *CronSpecUsecase) DeleteCronSpec(ctx context.Context, id uint64) error {
	param := map[string]interface{}{
		"deleted_at": time.Now(),
	}
	return uc.repo.UpdateCronSpec(ctx, id, param)
}

func (uc *CronSpecUsecase) ListCronSpec(ctx context.Context, f *CronSpecFilter) ([]*CronSpec, error) {
	return nil, nil
}

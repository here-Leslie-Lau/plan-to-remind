package biz

import "context"

type CronSpec struct {
	ID         uint64 `json:"id"`
	Desc       string `json:"desc"`
	Expression string `json:"expression"`
}

type CronSpecRepo interface {
	SaveCronSpec(ctx context.Context, cron *CronSpec) error
	GetCronSpec(ctx context.Context, id uint64) (*CronSpec, error)
	ListCronSpec(ctx context.Context, pageNum, pageSize int64) ([]*CronSpec, error)
	UpdateCronSpec(ctx context.Context, id uint64, params map[string]interface{}) error
}

type CronSpecUsecase struct {
	repo CronSpecRepo
}

func NewCronSpecUsecase(repo CronSpecRepo) *CronSpecUsecase {
	return &CronSpecUsecase{repo: repo}
}

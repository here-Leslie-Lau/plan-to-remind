package data

import (
	"context"
	"plan-to-remind/internal/biz"
)

var _ biz.CronSpecRepo = (*cronSpecRepo)(nil)

type cronSpecRepo struct {
	data *Data
}

func (c *cronSpecRepo) SaveCronSpec(ctx context.Context, cron *biz.CronSpec) error {
	panic("implement me")
}

func (c *cronSpecRepo) GetCronSpec(ctx context.Context, id uint64) (*biz.CronSpec, error) {
	panic("implement me")
}

func (c *cronSpecRepo) ListCronSpec(ctx context.Context, pageNum, pageSize int64) ([]*biz.CronSpec, error) {
	panic("implement me")
}

func (c *cronSpecRepo) UpdateCronSpec(ctx context.Context, id uint64, params map[string]interface{}) error {
	panic("implement me")
}

func NewCronSpecRepo(data *Data) biz.CronSpecRepo {
	return &cronSpecRepo{data: data}
}

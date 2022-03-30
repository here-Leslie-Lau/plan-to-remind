package data

import (
	"context"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/data/model"
)

var _ biz.CronSpecRepo = (*cronSpecRepo)(nil)

type cronSpecRepo struct {
	data *Data
}

func (c *cronSpecRepo) SaveCronSpec(ctx context.Context, cron *biz.CronSpec) error {
	res := &model.CronSpec{
		Desc:       cron.Desc,
		Expression: cron.Expression,
	}
	return c.data.db.Model(&model.CronSpec{}).Create(res).Error
}

func (c *cronSpecRepo) GetCronSpec(ctx context.Context, id uint64) (*biz.CronSpec, error) {
	panic("implement me")
}

func (c *cronSpecRepo) ListCronSpec(ctx context.Context, pageNum, pageSize int64) ([]*biz.CronSpec, error) {
	panic("implement me")
}

func (c *cronSpecRepo) UpdateCronSpec(ctx context.Context, id uint64, params map[string]interface{}) error {
	return c.data.db.Model(&model.CronSpec{}).Where("id = ?", id).Updates(params).Error
}

func NewCronSpecRepo(data *Data) biz.CronSpecRepo {
	return &cronSpecRepo{data: data}
}

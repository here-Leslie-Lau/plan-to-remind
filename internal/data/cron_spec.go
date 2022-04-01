package data

import (
	"context"
	"gorm.io/gorm"
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
	result := &model.CronSpec{}
	if err := c.data.db.Model(&model.CronSpec{}).Where("id = ?", id).First(result).Error; err != nil {
		return nil, err
	}
	return &biz.CronSpec{
		ID:         uint64(result.ID),
		Desc:       result.Desc,
		Expression: result.Expression,
	}, nil
}

func (c *cronSpecRepo) ListCronSpec(ctx context.Context, f *biz.CronSpecFilter) ([]*biz.CronSpec, error) {
	var list []*model.CronSpec
	if err := c.data.db.Scopes(limitCronSpecByFilter(f), limitOrderBy(f.OrderBy)).Find(&list).Error; err != nil {
		return nil, err
	}
	var result []*biz.CronSpec
	// maybe need to use go-copier
	for _, cronSpec := range list {
		result = append(result, &biz.CronSpec{
			ID:         uint64(cronSpec.ID),
			Desc:       cronSpec.Desc,
			Expression: cronSpec.Expression,
		})
	}
	return result, nil
}

func limitCronSpecByFilter(f *biz.CronSpecFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if f.OffSet > 0 && f.Limit > 0 {
			db = db.Offset(int(f.OffSet-1)).Limit(int(f.Limit))
		}
		return db
	}
}

func (c *cronSpecRepo) UpdateCronSpec(ctx context.Context, id uint64, params map[string]interface{}) error {
	return c.data.db.Model(&model.CronSpec{}).Where("id = ?", id).Updates(params).Error
}

func NewCronSpecRepo(data *Data) biz.CronSpecRepo {
	return &cronSpecRepo{data: data}
}

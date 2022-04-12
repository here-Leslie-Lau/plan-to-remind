package biz

import (
	"context"
	"time"
)

type CronSpec struct {
	ID         uint64 `json:"id"`
	Desc       string `json:"desc"`
	Expression string `json:"expression"`
	repo       CronSpecRepo
}

func NewCronSpec(ID uint64, desc string, expression string) *CronSpec {
	return &CronSpec{ID: ID, Desc: desc, Expression: expression, repo: RepoCron}
}

func NewDefaultCron(id uint64) *CronSpec {
	return &CronSpec{
		ID:   id,
		repo: RepoCron,
	}
}

func (c *CronSpec) Get(ctx context.Context) (*CronSpec, error) {
	return c.repo.GetCronSpec(ctx, c.ID)
}

func (c *CronSpec) Update(ctx context.Context) error {
	return c.repo.SaveCronSpec(ctx, c)
}

func (c *CronSpec) Create(ctx context.Context) error {
	return c.repo.SaveCronSpec(ctx, c)
}

func (c *CronSpec) Delete(ctx context.Context) error {
	param := map[string]interface{}{
		"deleted_at": time.Now(),
	}
	return c.repo.UpdateCronSpec(ctx, c.ID, param)
}

func (c *CronSpec) List(ctx context.Context, f *CronSpecFilter) ([]*CronSpec, error) {
	return c.repo.ListCronSpec(ctx, f)
}

type CronSpecFilter struct {
	OffSet  int64  `json:"off_set"`
	Limit   int64  `json:"limit"`
	OrderBy string `json:"order_by"`
}

type CronSpecRepo interface {
	SaveCronSpec(ctx context.Context, cron *CronSpec) error
	GetCronSpec(ctx context.Context, id uint64) (*CronSpec, error)
	ListCronSpec(ctx context.Context, f *CronSpecFilter) ([]*CronSpec, error)
	UpdateCronSpec(ctx context.Context, id uint64, params map[string]interface{}) error
}

var RepoCron CronSpecRepo

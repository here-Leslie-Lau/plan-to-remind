package data

import (
	"context"
	"gorm.io/gorm"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/data/model"
)

var _ biz.PlanRepo = (*PlanRepo)(nil)

const TableNamePlans = "plans"

// PlanRepo biz.PlanRepo实现，计划表仓库
type PlanRepo struct {
	data *Data
}

func (p *PlanRepo) CreatePlan(ctx context.Context, plan *biz.Plan) error {
	result := &model.Plan{
		State:    plan.State,
		Level:    plan.Level,
		CronID:   plan.CronId,
		DeadTime: plan.DeadTime,
		Name:     plan.Name,
	}
	return p.data.WithCtx(ctx).Model(&model.Plan{}).Create(result).Error
}

func (p *PlanRepo) GetPlan(ctx context.Context, id uint64) (*biz.Plan, error) {
	plan := &model.Plan{}
	if err := p.data.WithCtx(ctx).Model(&model.Plan{}).Where("id = ?", id).First(plan).Error; err != nil {
		return nil, err
	}
	result := &biz.Plan{
		ID:       uint64(plan.ID),
		State:    plan.State,
		Level:    plan.Level,
		CronId:   plan.CronID,
		DeadTime: plan.DeadTime,
		Name:     plan.Name,
	}
	// query cron desc
	var desc string
	err := p.data.WithCtx(ctx).Model(&model.CronSpec{}).Select("desc").Where("id = ?", id).Scan(&desc).Error
	if err != nil {
		return nil, err
	}
	result.CronDesc = desc
	return result, nil
}

func (p *PlanRepo) UpdatePlan(ctx context.Context, id uint64, param map[string]interface{}) error {
	return p.data.WithCtx(ctx).Model(&model.Plan{}).Where("id = ?", id).Updates(param).Error
}

// planModel 联表查询用到的结构体
type planModel struct {
	State    uint8  `json:"state"`
	Level    uint8  `json:"level"`
	ID       uint64 `json:"id"`
	CronID   uint64 `json:"cron_id"`
	DeadTime int64  `json:"dead_time"`
	Name     string `json:"name"`
	CronDesc string `json:"cron_desc"`
}

func (p *PlanRepo) ListPlanByFilter(ctx context.Context, f *biz.PlanFilter) ([]*biz.Plan, error) {
	var list []*planModel
	db := p.data.WithCtx(ctx).Select("plans.state AS state,plans.level AS level,plans.id AS id,plans.cron_id AS cron_id,plans.dead_time AS dead_time,plans.name AS name,cron_specs.desc AS cron_desc").
		Table(TableNamePlans).
		Joins("JOIN cron_specs ON plans.cron_id = cron_specs.id").Scopes(limitPlanByFilter(f), limitOrderBy(f.OrderBy))
	if err := db.Scan(&list).Error; err != nil {
		return nil, err
	}
	var result []*biz.Plan
	// maybe need to use go-copier
	for _, plan := range list {
		result = append(result, &biz.Plan{
			ID:       plan.ID,
			State:    plan.State,
			Level:    plan.Level,
			CronId:   plan.CronID,
			DeadTime: plan.DeadTime,
			Name:     plan.Name,
			CronDesc: plan.CronDesc,
		})
	}
	return result, nil
}

func limitPlanByFilter(f *biz.PlanFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if f.DeadTimeBegin > 0 {
			db = db.Where("dead_time >= ?", f.DeadTimeBegin)
		}
		if f.DeadTimeEnd > 0 {
			db = db.Where("dead_time <= ?", f.DeadTimeEnd)
		}
		if f.Offset > 0 && f.Limit > 0 {
			db = db.Offset(int(f.Offset - 1)).Limit(int(f.Limit))
		}
		return db
	}
}

func NewPlanRepo(data *Data) biz.PlanRepo {
	return &PlanRepo{data: data}
}

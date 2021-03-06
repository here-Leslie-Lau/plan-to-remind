package data

import (
	"context"
	"gorm.io/gorm"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/data/model"
	"time"
)

var _ biz.PlanRepo = (*PlanRepo)(nil)

const TableNamePlans = "plans"

// PlanRepo biz.PlanRepo实现，计划表仓库
type PlanRepo struct {
	data *Data
}

func (repo *PlanRepo) GetPlan(ctx context.Context, id uint64) (*biz.Plan, error) {
	plan := &model.Plan{}
	if err := repo.data.WithCtx(ctx).Model(&model.Plan{}).Where("id = ?", id).First(plan).Error; err != nil {
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
	return result, nil
}

func (repo *PlanRepo) DeletePlan(ctx context.Context, id uint64) error {
	return repo.data.WithCtx(ctx).Model(&model.Plan{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

// planModel 联表查询用到的结构体
type planModel struct {
	State      uint8  `json:"state"`
	Level      uint8  `json:"level"`
	ID         uint64 `json:"id"`
	CronID     uint64 `json:"cron_id"`
	DeadTime   int64  `json:"dead_time"`
	Name       string `json:"name"`
	CronDesc   string `json:"cron_desc"`
	Expression string `json:"expression"`
}

func (repo *PlanRepo) ListPlanByFilter(ctx context.Context, f *biz.PlanFilter) ([]*biz.Plan, error) {
	var list []*planModel
	db := repo.data.WithCtx(ctx).Select("plans.state AS state,plans.level AS level,plans.id AS id,plans.cron_id AS cron_id,plans.dead_time AS dead_time,plans.name AS name,cron_specs.desc AS cron_desc,cron_specs.expression AS expression").
		Table(TableNamePlans).
		Joins("JOIN cron_specs ON plans.cron_id = cron_specs.id").Scopes(limitPlanByFilter(f), limitOrderBy(f.OrderBy))
	if err := db.Scan(&list).Error; err != nil {
		return nil, err
	}
	var result []*biz.Plan
	// maybe need to use go-copier
	for _, plan := range list {
		result = append(result, &biz.Plan{
			ID:         plan.ID,
			State:      plan.State,
			Level:      plan.Level,
			CronId:     plan.CronID,
			DeadTime:   plan.DeadTime,
			Name:       plan.Name,
			CronDesc:   plan.CronDesc,
			Expression: plan.Expression,
		})
	}
	return result, nil
}

func limitPlanByFilter(f *biz.PlanFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if f.State != 0 {
			db = db.Where("state = ?", f.State)
		}
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

func (repo *PlanRepo) SavePlan(ctx context.Context, plan *biz.Plan) error {
	result := &model.Plan{
		State:    plan.State,
		Level:    plan.Level,
		CronID:   plan.CronId,
		DeadTime: plan.DeadTime,
		Name:     plan.Name,
	}
	sql := repo.data.WithCtx(ctx).Model(&model.Plan{})
	if plan.ID != 0 {
		// update
		return sql.Where("id = ?", plan.ID).Updates(result).Error
	}
	// create
	if err := sql.Create(result).Error; err != nil {
		return err
	}
	plan.ID = uint64(result.ID)
	return nil
}

func (repo *PlanRepo) SavePlanCompletion(ctx context.Context, completion *biz.PlanCompletion) error {
	plan := &model.PlanCompletion{
		TotalNums:     completion.TotalNums,
		CompletedNums: completion.CompletedNums,
		BeginAt:       completion.BeginAt,
		EndAt:         completion.EndAt,
		PlanId:        completion.PlanId,
	}
	sql := repo.data.WithCtx(ctx).Model(&model.PlanCompletion{})
	// update
	if err := sql.Where("plan_id = ? AND end_at >= ?", plan.PlanId, time.Now().Unix()).Updates(plan).Error; err != nil {
		return err
	}
	if sql.RowsAffected == 1 {
		return nil
	}
	// create
	return sql.Create(plan).Error
}

func limitCompletionByPlanId(planId uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("plan_id = ? AND end_at >= ?", planId, time.Now().Unix())
		return db
	}
}

func (repo *PlanRepo) GetLatestPlanCompletion(ctx context.Context, planId uint64) (*biz.PlanCompletion, error) {
	result := &model.PlanCompletion{}
	sql := repo.data.WithCtx(ctx).Model(&model.PlanCompletion{})
	sql = sql.Scopes(limitCompletionByPlanId(planId), limitOrderBy("id DESC"))
	if err := sql.First(result).Error; err != nil {
		return nil, err
	}
	return &biz.PlanCompletion{
		BeginAt:       result.BeginAt,
		EndAt:         result.EndAt,
		PlanId:        result.PlanId,
		TotalNums:     result.TotalNums,
		CompletedNums: result.CompletedNums,
	}, nil
}

func NewPlanRepo(data *Data) {
	biz.RepoPlan = &PlanRepo{data: data}
}

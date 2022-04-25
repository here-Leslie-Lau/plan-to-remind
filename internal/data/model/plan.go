package model

import "gorm.io/gorm"

// Plan 计划表
type Plan struct {
	gorm.Model
	// 1-on, 2-off
	State    uint8  `json:"state" gorm:"column:state; type:tinyint(2); comment:计划开关状态"`
	Level    uint8  `json:"level" gorm:"column:level; type:tinyint(8); comment:计划等级"`
	CronID   uint64 `json:"cron_id" gorm:"column:cron_id; comment:计划时间表ID"`
	DeadTime int64  `json:"dead_time" gorm:"column:dead_time; default:0; comment:计划失效时间"`
	Name     string `json:"name" gorm:"column:name; type:varchar(50); comment:计划名称"`
}

const (
	PlanStateOn = iota + 1
	PlanStateOff
)

// PlanCompletion 计划完成度
type PlanCompletion struct {
	gorm.Model
	TotalNums     int    `json:"total_nums" gorm:"column:total_nums; type:int(4); default:0; comment:应该完成的总次数"`
	CompletedNums int    `json:"completed_nums" gorm:"column:completed_nums; type:int(4); default:0; comment:完成的次数"`
	BeginAt       int64  `json:"begin_at" gorm:"column:begin_at; default:0;comment:统计开始时间"`
	EndAt         int64  `json:"end_at" gorm:"column:end_at; default:0; comment:统计结束时间"`
	PlanId        uint64 `json:"plan_id" gorm:"column:plan_id; not null; comment:计划ID"`
}

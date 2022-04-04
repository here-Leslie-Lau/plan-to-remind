package model

import "gorm.io/gorm"

// Plan 计划表
type Plan struct {
	gorm.Model
	State    uint8  `json:"state" gorm:"column:state; type:tinyint(2); comment:计划开关状态"`
	Level    uint8  `json:"level" gorm:"column:level; type:tinyint(8); comment:计划等级"`
	CronID   uint64 `json:"cron_id" gorm:"column:cron_id; comment:计划时间表ID"`
	DeadTime int64  `json:"dead_time" gorm:"column:dead_time; comment:计划失效时间"`
	Name     string `json:"name" gorm:"column:name; type:varchar(50); comment:计划名称"`
}

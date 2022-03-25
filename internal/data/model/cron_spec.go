package model

import (
	"gorm.io/gorm"
)

// CronSpec 计划时间表
type CronSpec struct {
	gorm.Model
	Desc       string `json:"desc" gorm:"column:desc; type:varchar(50); comment:计划时间描述"`
	Expression string `json:"expression" gorm:"column:expression; type:varchar(50); comment:cron表达式"`
}

func NewCronSpec(desc string) *CronSpec {
	return &CronSpec{Desc: desc}
}

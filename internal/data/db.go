package data

import (
	"gorm.io/gorm"
	"plan-to-remind/internal/conf"
	"plan-to-remind/internal/pkg/db"
)

func NewGormDb(data *conf.Data) (*gorm.DB, func()) {
	cfg := &db.DatabaseCfg{
		Username:     data.Database.Username,
		Password:     data.Database.Password,
		Url:          data.Database.Url,
		DatabaseName: data.Database.Database,
	}
	conn, f := db.NewGormConn(cfg)
	return conn.GetConn(), f
}

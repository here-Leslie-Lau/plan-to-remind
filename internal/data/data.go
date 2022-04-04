package data

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewGormDb, NewData, NewCronSpecRepo, NewPlanRepo)

// Data .
type Data struct {
	db *gorm.DB
}

func (d *Data) WithCtx(ctx context.Context) *gorm.DB {
	return d.db.WithContext(ctx)
}

func NewData(db *gorm.DB) (*Data, func(), error) {
	closeFunc := func() {
		// todo
		fmt.Println("closing data...")
	}
	return &Data{db: db}, closeFunc, nil
}

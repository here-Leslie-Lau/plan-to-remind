package data

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

// Data .
type Data struct {
	db *gorm.DB
}

func (d *Data) WithCtx(ctx context.Context) *gorm.DB {
	return d.db.WithContext(ctx)
}

func NewData(db *gorm.DB) (*Data, func(), error) {
	closeFunc := func() {
		fmt.Println("closing data...")
	}
	return &Data{db: db}, closeFunc, nil
}

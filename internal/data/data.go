package data

import (
	"context"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"gorm.io/gorm"
)

// Data .
type Data struct {
	db *gorm.DB
	// nacos client
	nacos config_client.IConfigClient
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

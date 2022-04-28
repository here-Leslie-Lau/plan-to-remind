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
	// redis client
	redis *RedisClient
}

func (d *Data) Redis() *RedisClient {
	return d.redis
}

func (d *Data) WithCtx(ctx context.Context) *gorm.DB {
	return d.db.WithContext(ctx)
}

func NewData(db *gorm.DB, nacos config_client.IConfigClient, redis *RedisClient) (*Data, func(), error) {
	closeFunc := func() {
		fmt.Println("closing data...")
		redis.Close()
	}
	return &Data{db: db, nacos: nacos, redis: redis}, closeFunc, nil
}

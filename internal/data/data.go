package data

import (
	"fmt"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewGormDb, NewData, NewCronSpecRepo)

// Data .
type Data struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) (*Data, func(), error) {
	closeFunc := func() {
		// todo
		fmt.Println("closing data...")
	}
	return &Data{db: db}, closeFunc, nil
}

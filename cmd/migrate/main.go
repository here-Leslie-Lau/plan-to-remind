package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"plan-to-remind/internal/conf"
	"plan-to-remind/internal/data/model"
	"plan-to-remind/internal/pkg/db"
)

var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "/home/hellotalk/go/src/plan-to-remind/configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	cfg := &db.DatabaseCfg{
		Username:     bc.Data.Database.Username,
		Password:     bc.Data.Database.Password,
		Url:          bc.Data.Database.Url,
		DatabaseName: "mysql",
	}
	conn, f := db.NewGormConn(cfg)
	defer f()

	// create database
	conn.CreateDatabase(bc.Data.Database.Database)
	// migrate
	cfg.DatabaseName = bc.Data.Database.Database
	dbConn, f2 := db.NewGormConn(cfg)
	defer f2()
	dbConn.AutoMigrate(model.Dst...)
}

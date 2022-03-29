package db

import (
	"plan-to-remind/internal/data/model"
	"testing"
)

func TestNewGormConn(t *testing.T) {
	cfg := &DatabaseCfg{
		Username:     "root",
		Password:     "1234",
		Url:          "127.0.0.1:3306",
		Port:         "3306",
		DatabaseName: "mysql",
	}
	_, f := NewGormConn(cfg)
	defer f()
}

func TestDropDb(t *testing.T) {
	cfg := &DatabaseCfg{
		Username:     "root",
		Password:     "1234",
		Url:          "127.0.0.1:3306",
		Port:         "3306",
		DatabaseName: "plan_to_remind",
	}
	conn, f := NewGormConn(cfg)
	defer f()

	if err := conn.DropDb(); err != nil {
		t.Fail()
	}
}

func TestAutoMigrate(t *testing.T) {
	cfg := &DatabaseCfg{
		Username:     "root",
		Password:     "1234",
		Url:          "127.0.0.1:3306",
		Port:         "3306",
		DatabaseName: "mysql",
	}
	conn, f := NewGormConn(cfg)
	defer f()
	// create database
	conn.cfg.DatabaseName = "demo"
	conn.CreateDatabase("")
	// migrate
	dst := []interface{}{
		model.CronSpec{},
	}
	cfg.DatabaseName = "demo"
	conn2, f2 := NewGormConn(cfg)
	defer f2()
	conn2.AutoMigrate(dst...)
}

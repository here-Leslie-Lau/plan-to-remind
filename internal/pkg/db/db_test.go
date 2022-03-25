package db

import "testing"

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
		DatabaseName: "demo",
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

	conn.cfg.DatabaseName = "demo"
	conn.AutoMigrate()
}

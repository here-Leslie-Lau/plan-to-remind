package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UtilsDbConn 目前持久化框架只考虑gorm实现
type UtilsDbConn struct {
	cfg  *DatabaseCfg
	conn *gorm.DB
}

func NewGormConn(cfg *DatabaseCfg) (*UtilsDbConn, func()) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Url, cfg.DatabaseName)
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
	// test ping
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := sqlDb.Ping(); err != nil {
		panic(fmt.Sprintf("mysql ping fail,%v", err))
	}
	cleanUp := func() {
		// no handler err
		_ = sqlDb.Close()
	}
	return &UtilsDbConn{
		cfg:  cfg,
		conn: db,
	}, cleanUp
}

func (u *UtilsDbConn) GetConn() *gorm.DB {
	return u.conn
}

func (u *UtilsDbConn) GetCfg() *DatabaseCfg {
	return u.cfg
}

func (u *UtilsDbConn) AutoMigrate(dst ...interface{}) {
	if err := u.conn.AutoMigrate(dst...); err != nil {
		panic(fmt.Sprintf("auto migrate tables failed,err:%v", err))
	}
	info := fmt.Sprintf("auto migrate tables database: %s succeed!", u.cfg.DatabaseName)
	fmt.Println(info)
}

func (u *UtilsDbConn) CreateDatabase(database string) {
	name := u.cfg.DatabaseName
	if database != "" {
		name = database
	}
	db := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARSET utf8mb4;", name)
	if err := u.conn.Exec(db).Error; err != nil {
		panic(fmt.Sprintf("create database failed,err:%v", err))
	}
	info := fmt.Sprintf("create database: %s succeed!", name)
	fmt.Println(info)
}

// DropDb Requests for the outside service to delete the database
func (u *UtilsDbConn) DropDb() error {
	db := "DROP DATABASE IF EXISTS " + u.cfg.DatabaseName + ";"
	if err := u.conn.Exec(db).Error; err != nil {
		panic("drop database failed, err:" + err.Error())
	}
	info := fmt.Sprintf("drop database: %s succeed!", u.cfg.DatabaseName)
	fmt.Println(info)
	return nil
}

package accessor

import (
	"sync"
	"time"

	"github.com/fast/internal/config"
	"github.com/fast/pkg/logger"
	"github.com/fast/pkg/mysql"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	onceDb   sync.Once
)

func InitializeDatabase() {

	cfg := config.Instance

	opts := []mysql.Option{
		mysql.WithMaxIdleConns(cfg.Mysql.MaxIdleConns),
		mysql.WithMaxOpenConns(cfg.Mysql.MaxOpenConns),
		mysql.WithConnMaxLifetime(time.Duration(cfg.Mysql.ConnMaxLifetime) * time.Minute),
	}

	if cfg.Mysql.EnableLog {
		opts = append(opts,
			mysql.WithLogging(logger.Get()),
			mysql.WithLogRequestIDKey("request_id"),
		)
	}
	if cfg.Server.EnableTrace {
		opts = append(opts, mysql.WithEnableTrace())
	}
	// setting mysql slave and master dsn addresses,
	// if there is no read/write separation, you can comment out the following piece of code
	opts = append(opts, mysql.WithRWSeparation(
		cfg.Mysql.SlavesDsn,
		cfg.Mysql.MastersDsn...,
	))

	// add custom gorm plugin
	//opts = append(opts, mysql.WithGormPlugin(yourPlugin))
	var err error
	database, err = mysql.Init(cfg.Mysql.Dsn, opts...)
	if err != nil {
		panic("mysql.Init error: " + err.Error())
	}
}

// CloseMysql close mysql
func CloseDatabase() error {
	if database == nil {
		return nil
	}

	sqlDB, err := database.DB()
	if err != nil {
		return err
	}
	if sqlDB != nil {
		return sqlDB.Close()
	}
	return nil
}

// GetDB get db
func Get() *gorm.DB {
	if database == nil {
		onceDb.Do(func() {
			InitializeDatabase()
		})
	}

	return database
}

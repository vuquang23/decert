package mysql

import (
	"decert/internal/pkg/config"
	"decert/internal/pkg/errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Instance() *gorm.DB {
	return db
}

func Init() error {
	if db != nil {
		return nil
	}

	dbCfg := config.Instance().DB
	gormDB, err := gorm.Open(
		mysql.Open(dbCfg.DNS()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(dbCfg.LogLevel)),
		},
	)
	if err != nil {
		return errors.NewInfraErrorDBConnect([]string{}, err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return errors.NewInfraErrorDBConnect([]string{}, err)
	}

	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(dbCfg.ConnLifeTime) * time.Second)

	if err = sqlDB.Ping(); err != nil {
		return errors.NewInfraErrorDBConnect([]string{}, err)
	}

	db = gormDB
	return nil
}

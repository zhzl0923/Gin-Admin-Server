package db

import (
	"fmt"
	"gin-admin/global"
	"gin-admin/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GetDialector func(dsn string) gorm.Dialector

var getDialectors = map[string]GetDialector{
	"mysql": mysql.Open,
}

func NewDB(dbSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbSetting.Username,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime,
	)
	dialector := getDialectors[dbSetting.DBType](dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.Config.Logger = logger.Default.LogMode(logger.Info)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dbSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbSetting.MaxOpenConns)
	return db, nil
}

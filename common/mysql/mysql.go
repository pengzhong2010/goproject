package mysql

import (
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(dsn string, idle, active int64, idleTimeout, execTimeout time.Duration) (db *gorm.DB, cf func(), err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxIdleTime(idleTimeout)
	sqlDB.SetConnMaxLifetime(execTimeout)
	sqlDB.SetMaxIdleConns(int(idle))
	sqlDB.SetMaxOpenConns(int(active))
	cf = func() { sqlDB.Close() }
	return
}

func ClearTable(tx *gorm.DB, table interface{}, now int64) (err error) {
	return tx.Model(table).Where("is_deleted=0").Where("sync_label!=?", now).Updates(map[string]interface{}{"is_deleted": now}).Error
}

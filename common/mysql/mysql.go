package mysql

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(dsn string, idle, active int64, idleTimeout, execTimeout time.Duration) (db *gorm.DB, cf func(), err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

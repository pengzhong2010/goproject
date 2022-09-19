package data

import (
	"goproject/app/thinktank/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewTodoRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (data *Data, cleanup func(), err error) {
	db, cf, err := NewMysql(c.Mysql)
	if err != nil {
		return
	}
	cleanup = func() {
		cf()
		log.NewHelper(logger).Info("closing the data resources")
	}
	data = &Data{
		Db: db,
	}
	return
}

func NewMysql(cfg *conf.Data_Mysql) (db *gorm.DB, cf func(), err error) {
	db, err = gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	//sqlDB.SetConnMaxIdleTime(cfg.IdleTimeout.AsDuration())
	//sqlDB.SetConnMaxLifetime(cfg.ExecTimeout.AsDuration())
	sqlDB.SetMaxIdleConns(int(cfg.Idle))
	sqlDB.SetMaxOpenConns(int(cfg.Active))
	//db.Logger = logger.Default.LogMode(logger.Info)
	cf = func() {}
	return
}

//func NewRedis() (r *redis.Redis, cf func(), err error) {
//	var (
//		cfg redis.Config
//		ct  paladin.Map
//	)
//	if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
//		return
//	}
//	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
//		return
//	}
//	r = redis.NewRedis(&cfg)
//	cf = func() { r.Close() }
//	return
//}

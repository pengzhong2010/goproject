package data

import (
	"goproject/common/redis"

	"goproject/app/thinktank/internal/conf"
	"goproject/common/mysql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewTodoRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Db    *gorm.DB
	Redis *redis.Redis
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (data *Data, cleanup func(), err error) {
	// 构造
	db, cf, err := mysql.NewMysql(c.Mysql.Dsn, c.Mysql.Idle, c.Mysql.Active, c.Mysql.IdleTimeout.AsDuration(), c.Mysql.ExecTimeout.AsDuration())
	if err != nil {
		return
	}
	redisDb, redisCf := redis.NewRedis(c.Redis.Network, c.Redis.Addr, c.Redis.Password, c.Redis.DbNum, c.Redis.Idle, c.Redis.Active, c.Redis.DialTimeout.AsDuration(), c.Redis.MaxConnLifetime.AsDuration())
	// 连接池
	data = &Data{
		Db:    db,
		Redis: redisDb,
	}
	// 析构
	cleanup = func() {
		cf()
		redisCf()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return
}

package data

import (
	"goproject/app/todo/internal/conf"
	"goproject/common/mysql"

	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
	NewMysqlClient,
	NewTodoRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, logger log.Logger) (*Data, func(), error) {

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: db,
	}, cleanup, nil
}
func NewMysqlClient(c *conf.Data) (db *gorm.DB, cf func(), err error) {
	db, cf, err = mysql.NewMysql(c.Mysql.Dsn, c.Mysql.Idle, c.Mysql.Active, c.Mysql.IdleTimeout.AsDuration(), c.Mysql.ExecTimeout.AsDuration())
	return
}

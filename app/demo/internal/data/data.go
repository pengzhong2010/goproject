package data

import (
	"context"

	v1 "goproject/api/helloworld/v1"
	"goproject/app/demo/internal/conf"
	"goproject/common/mysql"
	nc "goproject/common/nacos"
	"goproject/common/redis"
	"goproject/common/util"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewRegistrar,
	NewThinktankServiceClient,
	NewMysqlClient,
	NewRedisClient,
	NewTodoRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
	log   *log.Helper
	db    *gorm.DB
	redis *redis.Redis
	tc    v1.GreeterClient
}

// NewData .
func NewData(
	c *conf.Data,
	logger log.Logger,
	db *gorm.DB,
	redis *redis.Redis,
	tc v1.GreeterClient,
) (data *Data, cf func(), err error) {
	data = &Data{
		log:   log.NewHelper(log.With(logger, "module", "data")),
		db:    db,
		redis: redis,
		tc:    tc,
	}
	cf = func() {
		data.log.Info("closing the data resources")
	}
	return
}

func NewDiscovery(conf *conf.Registry) (r registry.Discovery, err error) {
	ncc := &nc.NacosConnConf{
		NacosAddr:        conf.Nacos.Addr,
		NacosPort:        conf.Nacos.Port,
		NacosNamespaceId: conf.Nacos.NamespaceId,
		NacosGroup:       conf.Nacos.Group,
	}
	r, err = ncc.Discovery()
	if err != nil {
		panic(err)
	}
	return
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	ncc := &nc.NacosConnConf{
		NacosAddr:        conf.Nacos.Addr,
		NacosPort:        conf.Nacos.Port,
		NacosNamespaceId: conf.Nacos.NamespaceId,
		NacosGroup:       conf.Nacos.Group,
	}
	r, err := ncc.Registry()
	if err != nil {
		panic(err)
	}
	return r
}

func NewThinktankServiceClient(ac *conf.Auth, r registry.Discovery) (c v1.GreeterClient, cf func(), err error) {
	env := util.GetEnv("ENV", "local")
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+"thinktank"+"."+env+".grpc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			jwt.Client(
				func(token *jwt2.Token) (interface{}, error) {
					return []byte(ac.Jwt.ServiceSecretKey), nil
				},
				jwt.WithSigningMethod(jwt2.SigningMethodHS256),
				jwt.WithClaims(func() jwt2.Claims {
					return &jwt2.MapClaims{}
				}),
			),
		),
	)
	if err != nil {
		panic(err)
	}
	c = v1.NewGreeterClient(conn)
	cf = func() { conn.Close() }
	return
}
func NewMysqlClient(c *conf.Data) (db *gorm.DB, cf func(), err error) {
	db, cf, err = mysql.NewMysql(c.Mysql.Dsn, c.Mysql.Idle, c.Mysql.Active, c.Mysql.IdleTimeout.AsDuration(), c.Mysql.ExecTimeout.AsDuration())
	return
}

func NewRedisClient(c *conf.Data) (db *redis.Redis, cf func(), err error) {
	db, cf = redis.NewRedis(c.Redis.Network, c.Redis.Addr, c.Redis.Password, c.Redis.DbNum, c.Redis.Idle, c.Redis.Active, c.Redis.DialTimeout.AsDuration(), c.Redis.MaxConnLifetime.AsDuration(), c.Redis.ReadTimeout.AsDuration(), c.Redis.WriteTimeout.AsDuration(), c.Redis.IdleTimeout.AsDuration())
	return
}

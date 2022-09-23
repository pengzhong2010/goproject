package main

import (
	"os"

	"goproject/app/demo/internal/conf"
	nc "goproject/common/nacos"
	"goproject/common/util"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	env, configFilePath string
	ncc                 *nc.NacosConnConf

	id, _ = os.Hostname()
)

func init() {
	initEnv()
}
func initEnv() {
	env = util.GetEnv("ENV", "local")
	configFilePath = util.GetEnv("CONFIG_FILE_PATH", "../../configs")
	if env != "local" {
		var err error
		ncc, err = nc.NewConfigConnConf()
		if err != nil {
			panic(err)
		}
	}
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, rr registry.Registrar) *kratos.App {
	if env == "local" {
		return kratos.New(
			kratos.ID(id),
			kratos.Name(Name),
			kratos.Version(Version),
			kratos.Metadata(map[string]string{}),
			kratos.Logger(logger),
			kratos.Server(
				gs,
				hs,
			),
		)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	// config
	cs, err := ncc.GetConfigSource(env, configFilePath)
	if err != nil {
		panic(err)
	}
	c := config.New(config.WithSource(cs...))
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}
	var ac conf.Auth
	if err := c.Scan(&ac); err != nil {
		panic(err)
	}
	Name = bc.Server.App.Name + "." + env
	Version = bc.Server.App.Version
	// di
	app, cleanup, err := wireApp(bc.Server, &rc, &ac, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

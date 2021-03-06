package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
	"plan-to-remind/internal/conf"
	"plan-to-remind/internal/data"
	"plan-to-remind/internal/pkg/limiter"
	"plan-to-remind/internal/server"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, cs *server.CronServer, cons *server.ConsumerServer) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
			cs,
			cons,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		// "service.name", Name,
		// "service.version", Version,
		// "trace.id", tracing.TraceID(),
		// "span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	closeFuncs := initRepo(bc.Data)
	defer func() {
		for _, closeFunc := range closeFuncs {
			closeFunc()
		}
	}()

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func initRepo(confData *conf.Data) []func() {
	db, f := data.NewGormDb(confData)
	nacos := data.NewConfigClient(confData)
	redis := data.NewRedisClient(confData)
	dataModel, f2, err := data.NewData(db, nacos, redis)
	if err != nil {
		panic(err)
	}
	data.NewCronSpecRepo(dataModel)
	data.NewPlanRepo(dataModel)
	data.NewConfigRepo(dataModel)

	limiter.InitPool(dataModel)
	return []func(){f, f2}
}

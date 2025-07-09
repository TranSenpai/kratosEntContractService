package main

import (
	"dormitory/internal/conf"
	"dormitory/internal/data"
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	_ "go.uber.org/automaxprocs"
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
	flag.StringVar(&flagconf, "conf", "configs/config.yaml", "config file path, e.g. -conf configs/config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, data *data.Data) (*kratos.App, error) {
	logHelper := log.NewHelper(log.With(logger, "module", "main"))
	if err := data.InitSchema(); err != nil {
		logHelper.Errorf("Schema init failed: %v", err)
	}
	app := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(gs),
	)

	return app, nil
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	// Load config from ./configs/config.yaml
	c := config.New(
		// WithSource specify the path of config file
		// flagconf is local variable store path in int method
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}
	// conf.Bootsrap is a struct generated from conf.proto in internal/conf,
	// including HTTP and gRPC, db
	var bc conf.Bootstrap
	// Load all config from yml assign to struct bc
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// Call wireApp to create all obj that each layer needed
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

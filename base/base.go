package base

import (
	"net"
	"path/filepath"
	"simple-k8/go-common/api-base"
	"strconv"

	"github.com/kataras/iris"
	"github.com/natefinch/lumberjack"
	"google.golang.org/grpc"
	slog "simple-k8/log"
)

var (
	VERSION         = "dev-snapshot"
	_RPCSERVER      *grpc.Server
	_SYSTEM_FAIL    = make(chan SystemFailure)
	RPC_SERVER_PORT int
	RPC_USE_TLS     bool
	API_HOST        string = "localhost"
	API_PORT        int    = 8889
)

func ConfigureProductVersion(v string) {
	VERSION = v
}

func ConfigureApiServer(host string, port int, root *apibase.Route, restrictSchema bool) error {
	API_HOST = host
	API_PORT = port
	app := iris.New()
	apibase.RegisterUUIDStringMacro(app)
	app.AttachLogger(&lumberjack.Logger{
		Filename:   filepath.Join(slog.LOGDIR, "api.log"),
		MaxSize:    slog.LOGGER_MAX_SIZE,
		MaxBackups: slog.LOGGER_MAX_BKS,
		MaxAge:     slog.LOGGER_MAX_AGE,
	})
	//app.StaticServe("easyagent", "/easyagent")
	if err := apibase.InitSchema(app, root, restrictSchema); err != nil {
		return err
	}

	go func() {
		err := app.Run(iris.Addr(net.JoinHostPort(host, strconv.Itoa(port))))
		if err != nil {
			SystemExitWithFailure(NETWORK_FAILURE, "API server failure: %v", err)
		}
	}()
	return nil
}

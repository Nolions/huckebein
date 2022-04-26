package main

import (
	"context"
	"errors"
	"flag"
	"github.com/nolions/huckebein/conf"
	"github.com/nolions/huckebein/internal/notify"
	"github.com/nolions/huckebein/internal/server"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	confPath string // config path
	logger   *zap.SugaredLogger
)

func main() {
	flag.StringVar(&confPath, "c", "conf.yaml", "default config path")
	flag.Parse()

	config, err := conf.New(confPath)
	if err != nil {
		log.Fatal(errors.New(err.Error()))
	}

	logger = zap.NewExample().Sugar()
	defer logger.Sync()

	ctx := context.Background()
	f := notify.NewsFirebase(ctx, logger)
	app := server.New(ctx, f, logger)
	serv := server.NewHttpServer(app, &config.HttpServ)
	serv.Run()
	shutdown(&config.App, serv)
}

func shutdown(conf *conf.App, srv *server.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	logger.Infof("get a signal %s. (%s) Server is shutting down ...\n", s.String(), conf.Name)

	// close http server with timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("http server shutdown: %v\n", err)
	}

	logger.Info("(%s) Server is exit.\n", conf.Name)
}

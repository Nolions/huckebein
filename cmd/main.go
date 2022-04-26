package main

import (
	"context"
	"errors"
	"flag"
	"github.com/nolions/huckebein/conf"
	"github.com/nolions/huckebein/internal/notify"
	"github.com/nolions/huckebein/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	confPath string // config path
)

func main() {
	flag.StringVar(&confPath, "c", "conf.yaml", "default config path")
	flag.Parse()

	config, err := conf.New(confPath)
	if err != nil {
		log.Fatal(errors.New(err.Error()))
	}

	ctx := context.Background()
	f := notify.NewsFirebase(ctx)
	app := server.New(ctx, f)
	serv := server.NewHttpServer(app, &config.HttpServ)
	serv.Run()
	shutdown(&config.App, serv)
}

func shutdown(conf *conf.App, srv *server.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	log.Printf("signal is %s", s)
	//log.Info().Msgf("get a signal %s. (%s) Server is shutting down ...", s.String(), project)

	// close http server with timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		//log.Fatal().Msgf("http server shutdown: %v", err)
	}

	//log.Info().Msgf("(%s) Server is exit.", project)
}

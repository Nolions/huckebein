package server

import (
	"context"
	"fmt"
	"github.com/nolions/huckebein/conf"
	"github.com/nolions/huckebein/internal/notify"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Ctx    context.Context
	Notify notify.Notify
	Logger *zap.SugaredLogger
}

type Server struct {
	HttpServer *http.Server
	Logger     *zap.SugaredLogger
}

func New(ctx context.Context, n *notify.Firebase, logger *zap.SugaredLogger) *Application {
	return &Application{
		Ctx:    ctx,
		Notify: n,
		Logger: logger,
	}
}

// NewHttpServer
// init http server
func NewHttpServer(app *Application, conf *conf.HttpServ) *Server {
	e := engine()
	app.router(e)

	addr := fmt.Sprintf(":%s", conf.Addr)
	app.Logger.Infof("Listening on %s\n", addr)
	h := &http.Server{
		Addr:         addr,
		Handler:      e,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &Server{
		HttpServer: h,
		Logger:     app.Logger,
	}
}

// Run
// run http server
func (app *Server) Run() {
	if err := app.HttpServer.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

// Shutdown
// shut down service
func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.Shutdown(ctx)
}

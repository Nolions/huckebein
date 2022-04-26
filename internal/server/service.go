package server

import (
	"context"
	"fmt"
	"github.com/nolions/huckebein/conf"
	"github.com/nolions/huckebein/internal/notify"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Ctx context.Context
	//Firebase *firebase.App
	Notify notify.Notify
}

type Server struct {
	HttpServer *http.Server
}

func New(ctx context.Context, n *notify.Firebase) *Application {
	return &Application{
		Ctx:    ctx,
		Notify: n,
	}
}

// NewHttpServer
// init http server
func NewHttpServer(app *Application, conf *conf.HttpServ) *Server {
	e := engine()
	app.router(e)

	addr := fmt.Sprintf(":%s", conf.Addr)
	log.Printf("Listening on %s", addr)
	h := &http.Server{
		Addr:         addr,
		Handler:      e,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &Server{
		HttpServer: h,
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

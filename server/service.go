package server

import (
	"context"
	"github.com/nolions/huckebein/notify"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	Ctx context.Context
	//Firebase *firebase.App
	Notify notify.Notify
}

type Server struct {
	httpServer *http.Server
}

func New(ctx context.Context, n *notify.Firebase) *Application {
	return &Application{
		Ctx:    ctx,
		Notify: n,
	}
}

// NewHttpServer
// init http server
func NewHttpServer(app *Application) *Server {
	e := engine()
	app.router(e)

	log.Printf("Listening on %s", ":7777")
	h := &http.Server{
		Addr:         ":7777",
		Handler:      e,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &Server{
		httpServer: h,
	}
}

// Run
// run http server
func (app *Server) Run() {
	if err := app.httpServer.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (app *Server) SignalProcess() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	s := <-c
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		log.Printf("signal is %s", s)
		app.httpServer.Close()
		return
	case syscall.SIGHUP:
	default:
		return
	}
}

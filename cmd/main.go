package main

import (
	"context"
	"github.com/nolions/huckebein/internal/notify"
	"github.com/nolions/huckebein/internal/server"
)

func main() {
	ctx := context.Background()

	f := notify.NewsFirebase(ctx)
	serv := server.New(ctx, f)
	app := server.NewHttpServer(serv)
	app.Run()
	go app.SignalProcess()
}

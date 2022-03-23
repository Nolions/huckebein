package main

import (
	"context"
	"github.com/nolions/huckebein/notify"
	"github.com/nolions/huckebein/server"
)

func main() {
	ctx := context.Background()

	f := notify.NewsFirebase(ctx)
	serv := server.New(ctx, f)
	app := server.NewHttpServer(serv)
	go app.SignalProcess()
	app.Run()
}

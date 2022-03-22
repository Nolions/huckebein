package main

import (
	"context"
	"github.com/nolions/huckebein/server"
)

func main() {
	ctx := context.Background()

	serv := server.New(ctx, server.NewsFirebase(ctx))

	app := server.NewHttpServer(serv)
	go app.SignalProcess()
	app.Run()
}

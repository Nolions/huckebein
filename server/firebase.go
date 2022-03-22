package server

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"log"
)

// NewsFirebase
// init Firebase app
func NewsFirebase(ctx context.Context) *firebase.App {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

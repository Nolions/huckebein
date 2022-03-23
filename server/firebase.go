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

type Metadata map[string]interface{}

type NotifyMsg struct {
	DeviceToke string   `json:"device_toke" validate:"required, string"`
	Title      string   `json:"title" validate:"required, string"`
	Message    string   `json:"message"  validate:"required, string"`
	Metadata       Metadata `json:"metadata"`
}

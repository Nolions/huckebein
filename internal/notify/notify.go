package notify

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/nolions/huckebein/internal/model"
	"log"
)

type Notify interface {
	SendNotify(*model.NotifyReq)
	SendMultiNotify(*model.MultiNotifyReq)
	BatchSendNotify([]model.NotifyReq)
}

type Firebase struct {
	Ctx      context.Context
	Firebase *firebase.App
}

// NewsFirebase
// init Firebase
func NewsFirebase(ctx context.Context) *Firebase {
	return &Firebase{
		Ctx:      ctx,
		Firebase: newsFirebaseApp(ctx),
	}
}

// NewsFirebaseApp
// init Firebase app
func newsFirebaseApp(ctx context.Context) *firebase.App {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

func (f *Firebase) SendNotify(msg *model.NotifyReq) {
	client, err := f.Firebase.Messaging(f.Ctx)
	if err != nil {
		log.Printf("error getting Messaging client: %v\n", err)
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: msg.Title,
			Body:  msg.Message,
		},
		Token: msg.DeviceToke,
	}

	response, err := client.Send(f.Ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully sent message:%v", response)
}

// SendMultiNotify
// Send a msg to multi device
func (f *Firebase) SendMultiNotify(msg *model.MultiNotifyReq) {
	client, err := f.Firebase.Messaging(f.Ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// See documentation on defining a message payload.
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: msg.Title,
			Body:  msg.Message,
		},
		Tokens: msg.DeviceTokes,
	}

	response, err := client.SendMulticast(f.Ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully sent message:%v", response)
}

// BatchSendNotify
// Batch send message
func (f *Firebase) BatchSendNotify(msgs []model.NotifyReq) {
	client, err := f.Firebase.Messaging(f.Ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	var messages []*messaging.Message
	for _, msg := range msgs {
		messages = append(messages,
			&messaging.Message{
				Notification: &messaging.Notification{
					Title: msg.Title,
					Body:  msg.Message,
				},
				Token: msg.DeviceToke,
			},
		)
	}

	resp, err := client.SendAll(context.Background(), messages)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(" Successfully batch send notification message:%v", resp)
}

package notify

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/nolions/huckebein/internal/model"
	"go.uber.org/zap"
)

type Notify interface {
	SendNotify(*model.NotifyReq)
	SendMultiNotify(*model.MultiNotifyReq)
	BatchSendNotify([]model.NotifyReq)
}

type Firebase struct {
	Ctx      context.Context
	Firebase *firebase.App
	Logger   *zap.SugaredLogger
}

// NewsFirebase
// init Firebase
func NewsFirebase(ctx context.Context, logger *zap.SugaredLogger) *Firebase {
	return &Firebase{
		Ctx:      ctx,
		Firebase: newsFirebaseApp(ctx, logger),
		Logger:   logger,
	}
}

// NewsFirebaseApp
// init Firebase app
func newsFirebaseApp(ctx context.Context, logger *zap.SugaredLogger) *firebase.App {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		logger.Errorf("error initializing app: %v\n", err)
	}

	return app
}

func (f *Firebase) SendNotify(msg *model.NotifyReq) {
	client, err := f.Firebase.Messaging(f.Ctx)
	if err != nil {
		f.Logger.Errorf("error getting Messaging client: %v\n", err)
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
		f.Logger.Errorf("send Messaging error: %v\n", err)
	}
	f.Logger.Infof("Successfully sent message:%v\n", response)
}

// SendMultiNotify
// Send a msg to multi device
func (f *Firebase) SendMultiNotify(msg *model.MultiNotifyReq) {
	client, err := f.Firebase.Messaging(f.Ctx)
	if err != nil {
		f.Logger.Errorf("error getting Messaging client: %v\n", err)
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
		f.Logger.Errorf("send Multicast Messaging error: %v\n", err)
	}
	f.Logger.Infof("Successfully sent message:%v\n", response)
}

// BatchSendNotify
// Batch send message
func (f *Firebase) BatchSendNotify(msgs []model.NotifyReq) {
	client, err := f.Firebase.Messaging(f.Ctx)
	if err != nil {
		f.Logger.Errorf("error getting Messaging client: %v\n", err)
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
		f.Logger.Errorf("send Batch Messaging error: %v\n", err)
	}
	f.Logger.Infof("Successfully sent message:%v\n", resp)
}

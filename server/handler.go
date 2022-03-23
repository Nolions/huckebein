package server

import (
	"firebase.google.com/go/v4/messaging"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (serv Application) handler(e *gin.Engine) {
	e.GET("/health", serv.healthHandler)
	e.POST("/notification", serv.notificationHandler)
}

func (serv Application) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (serv Application) notificationHandler(c *gin.Context) {
	client, err := serv.Firebase.Messaging(serv.Ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	notifyMsg := &NotifyMsg{}
	err = c.BindJSON(&notifyMsg)
	log.Printf("req:%v", notifyMsg)
	if err != nil {
		log.Fatalf("request data error: %v\n", err)
		return
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: notifyMsg.Title,
			Body:  notifyMsg.Message,
		},
		Token: notifyMsg.DeviceToke,
	}

	response, err := client.Send(serv.Ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully sent message:%v", response)

	c.Status(http.StatusNoContent)
}

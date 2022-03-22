package server

import (
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (serv Application) handler(e *gin.Engine) {
	e.GET("/health", serv.healthHandler)
	e.GET("/notification", serv.notificationHandler)
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

	registrationToken := "c2b4uXJlRdGxGjh_po0LU0:APA91bGyjUpMmb7WBcgd5CBer8RCu1FXnZAomcjCzOluF2dEZnWdPKzLH_3ACuEPwuRqvbfLpyyjiIx1grxvvMrHxHDOjvoQofYZpnNaC-l_8RR0TudWq4BzTI9vWe4HM9wYnYQjVJ6M"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"ssss": "850",
			"ss":   "2:45",
		},
		Token: registrationToken,
	}

	response, err := client.Send(serv.Ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

	c.Status(http.StatusNoContent)
}

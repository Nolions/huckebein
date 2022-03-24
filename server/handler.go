package server

import (
	"github.com/nolions/huckebein/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (serv Application) router(e *gin.Engine) {
	e.NoMethod(HandleNotFound)
	e.NoRoute(HandleNotFound)

	e.GET("/health", ErrHandler(serv.healthHandler))
	e.POST("/notify", ErrHandler(serv.notifyHandler))
	e.POST("/multiNotify", ErrHandler(serv.multiNotifyHandler))
}

func (serv Application) healthHandler(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

	return nil
}

func (serv Application) notifyHandler(c *gin.Context) error {
	notifyMsg := &model.NotifyReq{}
	err := c.BindJSON(&notifyMsg)
	log.Printf("req:%v", notifyMsg)
	if err != nil {
		log.Fatalf("request data error: %v\n", err)
		return err
	}

	serv.Notify.SendNotify(notifyMsg)

	c.Status(http.StatusNoContent)

	return nil
}

func (serv Application) multiNotifyHandler(c *gin.Context) error {
	json := &model.MultiNotifyReq{}
	err := c.BindJSON(&json)
	log.Printf("req:%v", json)
	if err != nil {
		log.Fatalf("request data error: %v\n", err)
		return err
	}

	serv.Notify.SendMultiNotify(json)

	return nil
}

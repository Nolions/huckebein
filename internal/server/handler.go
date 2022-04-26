package server

import (
	"github.com/nolions/huckebein/internal/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app Application) router(e *gin.Engine) {
	e.HandleMethodNotAllowed = true
	e.NoMethod(HandleNoAllowMethod)
	e.NoRoute(HandleNotFound)

	e.GET("/health", ErrHandler(app.healthHandler))
	e.POST("/notify", ErrHandler(app.notifyHandler))
	e.POST("/multiNotify", ErrHandler(app.multiNotifyHandler))
	e.POST("/batchNotify", ErrHandler(app.batchNotifyHandler))
}

func (app Application) healthHandler(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

	return nil
}

func (app Application) notifyHandler(c *gin.Context) error {
	notifyMsg := &model.NotifyReq{}
	err := c.BindJSON(&notifyMsg)
	app.Logger.Info("req:%v\n", notifyMsg)
	if err != nil {
		app.Logger.Errorf("request data error: %v\n", err)
		return err
	}

	app.Notify.SendNotify(notifyMsg)

	c.Status(http.StatusNoContent)

	return nil
}

func (app Application) multiNotifyHandler(c *gin.Context) error {
	json := &model.MultiNotifyReq{}
	err := c.BindJSON(&json)
	log.Printf("req:%v", json)
	if err != nil {
		app.Logger.Errorf("request data error: %v\n", err)
		return err
	}

	app.Notify.SendMultiNotify(json)

	return nil
}

func (app Application) batchNotifyHandler(c *gin.Context) error {
	var json []model.NotifyReq
	err := c.BindJSON(&json)
	log.Printf("req:%v", json)
	if err != nil {
		app.Logger.Errorf("request data error: %v\n", err)
		return err
	}

	app.Notify.BatchSendNotify(json)

	return nil
}

package server

import (
	"github.com/nolions/huckebein/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (serv Application) router(e *gin.Engine) {
	e.GET("/health", serv.healthHandler)
	e.POST("/notification", serv.notifyHandler)
}

func (serv Application) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (serv Application) notifyHandler(c *gin.Context) {
	notifyMsg := &model.NotifyReq{}
	err := c.BindJSON(&notifyMsg)
	log.Printf("req:%v", notifyMsg)
	if err != nil {
		log.Fatalf("request data error: %v\n", err)
		return
	}

	serv.Notify.SendNotify(notifyMsg)

	c.Status(http.StatusNoContent)
}

package handler

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:" status"`
}

func newErrorResponse(c *gin.Context, stausCode int, message string) {
	log.Error(message)
	c.AbortWithStatusJSON(stausCode, errorResponse{message})
}

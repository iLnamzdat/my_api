package handler

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	massage string `json:"message"`
}

func newErrorResponse(c *gin.Context, stausCode int, message string) {
	log.Error(message)

	c.AbortWithStatusJSON(stausCode, errorResponse{message})
}

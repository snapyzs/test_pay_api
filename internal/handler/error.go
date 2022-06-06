package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponce struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func newErrResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponce{message, statusCode})
}

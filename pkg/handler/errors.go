package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, response{Message: message})
}

package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type error struct{
	Message string `json:"message"`
}

func newErrorMessage(c *gin.Context, statusCode int, message string){
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
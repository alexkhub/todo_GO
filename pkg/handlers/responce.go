package handlers

import (
	"log"
	todogo "todo"

	"github.com/gin-gonic/gin"
)

type errorResponce struct{
	Message string `json:"message"`
}

type getListTakskResponse struct {
	Data []todogo.ListTask `json:"data"`
}

func newErrorMessage(c *gin.Context, statusCode int, message string){
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponce{message})
}
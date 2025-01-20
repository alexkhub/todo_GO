package handlers

import (
	"errors"
	"strconv"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const(
	signingKey = "fdljdcsdcsv232e3cdjif"
)


func (h *Handler) parseAuthHeader(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorMessage(c, http.StatusForbidden, errors.New("empty auth header").Error())
		return 
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorMessage(c, http.StatusForbidden, errors.New("invalid auth header").Error())
		return 
	}
	
	if len(headerParts[1]) == 0 {
		
		newErrorMessage(c, http.StatusForbidden, errors.New("token is empty").Error())
		return 
	}
	res, err := h.auth.Parse(headerParts[1])

	if err != nil {
		newErrorMessage(c, http.StatusForbidden,  err.Error())
		return 
	}
	c.Set("user_id", res)
	
	
}

func getUserId(c *gin.Context) (int, error){

	id, ok := c.Get("user_id")
	if !ok {

		return 0, errors.New("user id not found")
	}
	
	idInt, err := strconv.Atoi(id.(string))

	if err != nil {
		return 0, errors.New("user id not convert")
	}
	return idInt, nil
}
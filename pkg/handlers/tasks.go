package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) TaskList(c *gin.Context) {
	res, _ := c.Get("user_id")
	c.JSON(200, gin.H{
		
		"user": res,
	})
}

func (h *Handler) TaskDetail(c *gin.Context) {

}

func (h *Handler) TaskUpdate(c *gin.Context) {
	
}

func (h *Handler) TaskDelete(c *gin.Context) {

}

func (h *Handler) TaskCreate(c *gin.Context){
	
}
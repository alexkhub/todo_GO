package handlers

import (
	"net/http"
	"strconv"
	todogo "todo"

	v "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (h *Handler) TaskList(c *gin.Context) {
	id, err := getUserId(c)

	if err!= nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return 
    }
    list, err := h.servies.TaskList(id)

    if err != nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return 
	}
    c.JSON(http.StatusOK, getListTakskResponse{
        Data: list,
    })

    
}

func (h *Handler) TaskDetail(c *gin.Context) {
    user_id, err := getUserId(c)

	if err!= nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return 
    }
    task_id, err := strconv.Atoi(c.Param("id"))
    if err!= nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return 
    }
    task, err := h.servies.TaskDetail(user_id, task_id)

    if err!=nil{
        newErrorMessage(c, http.StatusNotFound, err.Error())
        return
    }

    c.JSON(http.StatusOK, task)
}

func (h *Handler) TaskUpdate(c *gin.Context) {
    var input todogo.UpdateTask
    user_id, err := getUserId(c)
    task_id, err := strconv.Atoi(c.Param("id"))
    if err!= nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return 
    }

	if err!= nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return 
    }
    
    if err:= c.BindJSON(&input); err != nil{ 
        newErrorMessage(c, http.StatusBadRequest, err.Error())   
        return 
    }
    _, err = v.ValidateStruct(input)
    if err!= nil{
        newErrorMessage(c, http.StatusBadRequest, err.Error())
        return 
    }
    task, err := h.servies.TaskUpdate(user_id, task_id, input)
    if err!=nil{
        newErrorMessage(c, http.StatusNotFound, err.Error())
        return
    }

    c.JSON(http.StatusOK, task)
	
}

func (h *Handler) TaskDelete(c *gin.Context) {
    user_id, err := getUserId(c)

	if err!= nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return 
    }
    task_id, err := strconv.Atoi(c.Param("id"))
    if err!= nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return 
    }
    err = h.servies.TaskDelete(user_id, task_id)
    if err != nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusNoContent, gin.H{})

}

func (h *Handler) TaskCreate(c *gin.Context){
	var input todogo.CreateTask
	
	
	id, err := getUserId(c)

	if err!= nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return 
    }
    
    if err:= c.BindJSON(&input); err != nil{ 
        newErrorMessage(c, http.StatusBadRequest, err.Error())   
        return 
    }
    _, err = v.ValidateStruct(input)
    if err!= nil{
        newErrorMessage(c, http.StatusBadRequest, err.Error())
        return 
    }
	object, err := h.servies.Task.CreateTask(id, input)

	if err != nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return 
	}

	c.JSON(http.StatusOK, gin.H{
        "id" : object ,
    })
}
package handlers

import (
	"net/http"
	"todo"
	"github.com/gin-gonic/gin"
     v "github.com/asaskevich/govalidator"

)




func (h *Handler) MainPage(c * gin.Context){
    c.JSON(200, gin.H{
        "message": "pong",
    })
}


func (h *Handler) Register(c *gin.Context){
    var input todogo.RegisterUser
    
    if err:= c.BindJSON(&input); err != nil{
        
        newErrorMessage(c, http.StatusBadRequest, err.Error())
        
        return 
    }
    _, err := v.ValidateStruct(input)
    if err!= nil{
        newErrorMessage(c, http.StatusBadRequest, err.Error())
        
        return 
    }
    
    id, err :=  h.servies.Authorization.CreateUser(input)
    
    if err != nil{
        newErrorMessage(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "id" : id,
    })
    
}

func (h *Handler) Login(c *gin.Context){
    var input todogo.LoginUser
    
    if err:= c.BindJSON(&input); err != nil{
        
        newErrorMessage(c, http.StatusBadRequest, err.Error())
        
        return 
    }
    _, err := v.ValidateStruct(input)
    if err!= nil{
        newErrorMessage(c, http.StatusBadRequest, err.Error())
        
        return 
    }
    
    token, err :=  h.servies.Authorization.LoginUser(input)
    
    if err != nil{
        newErrorMessage(c, http.StatusUnauthorized, err.Error())
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "token" : token,
    })
}



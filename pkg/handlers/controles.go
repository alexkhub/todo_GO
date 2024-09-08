package handlers

import (
	
	"github.com/gin-gonic/gin"
	"net/http"	
	"database/sql"
)

type User struct{
    Id int `json:"id"`
    Username string `json:"username"`
    Email string `json:"email"`
}

type Handler_DB struct{
    DB  *sql.DB
}


func (h *Handler_DB) AllUser(c *gin.Context){
    users := []*User{}
    rows, err := h.DB.Query("Select id, username, email from users")
    if err != nil{
        panic(err)
    }
    for rows.Next(){
        user := &User{}
        err = rows.Scan(&user.Id, &user.Username, &user.Email)
        if err != nil{
            return 
        }
        users = append(users, user )
        
        

    }
    defer rows.Close()
	
	c.IndentedJSON(http.StatusOK, users)
}





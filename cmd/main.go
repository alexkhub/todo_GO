package main

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

)

type Handler_DB struct{
    DB  *sql.DB
}
type User struct{
    Id int `json:"id"`
    Username string `json:"username"`
    Email string `json:"email"`
}



func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("NAME"))
	fmt.Print(psqlInfo)
	psqlInfo = "host=localhost port=5432 user=postgres password=34583458Ak dbname=todo_GO sslmode=disable"
	db, _ := sql.Open("postgres", psqlInfo)
	db.SetMaxIdleConns(10)

	err := db.Ping()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("DB CONNECT")
	}
	defer db.Close()

	


	

	router := gin.Default()
    

    router.Run("localhost:8080")
	fmt.Println("Server start")

}


	
	

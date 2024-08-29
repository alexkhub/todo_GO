package main

import (
	"database/sql"
	"fmt"

	"os"
	"github.com/gin-gonic/gin"
	hd "todo/pkg/handlers"

	_ "github.com/lib/pq"
)



func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("NAME"))
	fmt.Print(psqlInfo)
	psqlInfo = "host=localhost port=5432 user=postgres password=root dbname=todo_GO sslmode=disable"
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
    router.GET("/albums",hd.Alltasks)

    router.Run("localhost:8080")


    router.Run("localhost:8080")

	
	fmt.Println("Server start")

}

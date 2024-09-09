package main

import (
	"database/sql"
	"fmt"
	"log"
	todogo "todo"
	"todo/pkg/handlers"

	_ "github.com/lib/pq"
)


type User struct{
    Id int `json:"id"`
    Username string `json:"username"`
    Email string `json:"email"`
}



func main() {
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    // "password=%s dbname=%s sslmode=disable",
    // os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("NAME"))
	
	psqlInfo := "host=localhost port=5432 user=postgres password=root dbname=todo_GO sslmode=disable"
	db, _ := sql.Open("postgres", psqlInfo)
	db.SetMaxIdleConns(10)

	err := db.Ping()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("DB CONNECT")
	}
	defer db.Close()

	


	
	my_handlers := new(handlers.Handler)
	srv := new(todogo.Server)
	if err := srv.Run(my_handlers.InitRouter()); err != nil {
		log.Fatalf("server dont start")
	} 

}


	
	

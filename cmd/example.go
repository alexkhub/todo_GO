package main

import (
	"log"
	todogo "todo"
	"todo/pkg/handlers"
	"todo/pkg/repository"
	"todo/pkg/service"

	_ "github.com/lib/pq"
	

)



func main() {
	
	db, err := repository.NewDBConnect()
	

	
	if err != nil{
		log.Fatalln("db err")
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)



	my_handlers := handlers.NewHandler(services)
	srv := new(todogo.Server)
	if err := srv.Run(my_handlers.InitRouter()); err != nil {
		log.Fatalf("server dont start")
	} 

}


  
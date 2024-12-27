package main

import (
	"log"
	todogo "todo"
	"todo/pkg/handlers"
	"todo/pkg/repository"
	"todo/pkg/service"

	_ "github.com/lib/pq"
	

)

const (
	signingKey = "fdljdcsdcsv232e3cdjif"
    signingKey2 = "fdvsgf34$%MJP&(^JGTOIOI)"
)


func main() {
	
	db, err := repository.NewDBConnect()
	

	
	if err != nil{
		log.Fatalln("db err")
	}

	repos := repository.NewRepository(db)
	auth := service.NewManager(signingKey, signingKey2)

	services := service.NewService(service.Deps{
		Repos: repos,
	})
	



	my_handlers := handlers.NewHandler(services, auth)
	srv := new(todogo.Server)
	if err := srv.Run(my_handlers.InitRouter()); err != nil {
		log.Fatalf("server dont start")
	} 

}


  
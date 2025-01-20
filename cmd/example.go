package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	go func() {
		if err := srv.Run(my_handlers.InitRouter()); err != nil {
			log.Fatalf("server dont start")
		} 
		}()

		log.Print("TodoApp Started")
	
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit
	
		log.Print("TodoApp Shutting Down")
	
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatal("error occured while running http server: %s", err.Error())
		}
	
		if err := db.Close(); err != nil {
			log.Fatal("error occured while running http server: %s", err.Error())
		}

}


  
package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)


func NewDBConnect() (*sqlx.DB, error){
	db, err :=  sqlx.Open("postgres", "host=localhost port=5432 user=root password=alex0000  dbname=my_db sslmode=disable")
	if err != nil{
		fmt.Println("DB ERROR")
		return nil, err 
	}
	
	err = db.Ping()
	if err != nil{
		fmt.Println("DB ERROR")
		return nil, err 
	}
	return db, nil
}
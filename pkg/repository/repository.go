package repository

import (
    "github.com/jmoiron/sqlx"
    todogo "todo"

)

const (
    UserTable = "users"
    TaskTable = "task"
)


type Authorization interface{
    CreateUser(user todogo.RegisterUser) (int, error)
    GetUser(user todogo.LoginUser) (int, string, error)

}

type Task interface{

}

type Repository struct{
    Authorization
    Task
}
func NewRepository(db *sqlx.DB) *Repository{
    return &Repository{
        Authorization: NewAuthPostgres(db),
        
    }
}
package repository

import (
    "github.com/jmoiron/sqlx"
    todogo "todo"

)

const (
    UserTable = "users"
    TaskTable = "tasks"
    JWTTable = "refresh"
)


type Authorization interface{
    CreateUser(user todogo.RegisterUser) (int, error)
    GetUser(user todogo.LoginUser) (int, string, error)
    CreateNewRefresh(hash_id float64, refresh_token string)(error)

}

type Task interface{
    CreateTask(user_id int, task todogo.CreateTask)(int, error)
    TaskList(user_id int)([]todogo.ListTask, error)
    TaskDetail(user_id int, task_id int)(todogo.ListTask, error)
    TaskDelete(user_id int, task_id int)( error)
    TaskUpdate(user_id int, task_id int, input todogo.UpdateTask)(todogo.ListTask, error)
}

type Repository struct{
    Authorization
    Task
}

func NewRepository(db *sqlx.DB) *Repository{
    return &Repository{
        Authorization: NewAuthPostgres(db),
        Task: NewTaskPostgres(db),
        
    }
}
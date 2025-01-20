package service

import (
	todogo "todo"
	"todo/pkg/repository"
)

const (
	signingKey = "fdljdcsdcsv232e3cdjif"
    signingKey2 = "fdvsgf34$%MJP&(^JGTOIOI)"
)

type Authorization interface{
    CreateUser(user todogo.RegisterUser) (int, error)
    LoginUser(user todogo.LoginUser) (todogo.JWTToken, error)
    RefreshJWT(refresh_token string) (todogo.JWTToken, error)
    
}

type JWTManager interface{
	CreateJwtAccess(user_id string) (string, error)
    CreateJwtRefresh(user_id string) (string, error)
    Parse(accessToken string) (string, error)    

}


type Task interface{
    CreateTask(user_id int, task todogo.CreateTask)(int, error)
    TaskList(user_id int)([]todogo.ListTask, error)
    TaskDetail(user_id int, task_id int)(todogo.ListTask, error)
    TaskDelete(user_id int, task_id int)( error)
    TaskUpdate(user_id int, task_id int, input todogo.UpdateTask)(todogo.ListTask, error)
}

type Deps struct {
    Repos *repository.Repository
    JWTManager JWTManager 
}

type Service struct{
    Authorization
    Task
    
}
func NewService(deps Deps) *Service{
    new_jwt_manager := NewManager(signingKey, signingKey2)
    new_auth_service := NewAuthService(deps.Repos.Authorization, new_jwt_manager)
    task := NewTaskService(deps.Repos.Task)
    
    
    return &Service{
        Authorization: new_auth_service,
        Task: task,
        

    }
}
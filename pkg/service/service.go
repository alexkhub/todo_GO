package service

import (
	todogo "todo"
	"todo/pkg/repository"
)

type Authorization interface{
    CreateUser(user todogo.RegisterUser) (int, error)
    LoginUser(user todogo.LoginUser) (string, error)
}

type Task interface{

}

type Service struct{
    Authorization
    Task
}
func NewService(repos *repository.Repository) *Service{
    return &Service{
        Authorization: NewAuthService(repos.Authorization),
        
    }
}
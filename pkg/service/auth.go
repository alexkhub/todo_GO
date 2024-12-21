package service

import (
	"fmt"
	todogo "todo"
	"todo/pkg/repository"
)


type AuthService struct{
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService{
	return &AuthService{repos:  repos}
}

func (s *AuthService) CreateUser (user todogo.RegisterUser) (int, error){
	user.Password, _ = HashPassword(user.Password)
	
	return s.repos.CreateUser(user)
}


func (s *AuthService) LoginUser (user todogo.LoginUser) (string, error){
	
	
	id, hash, err := s.repos.GetUser(user)
	
	if err!= nil{
		return "", err
	}
	chech := CheckPasswordHash(user.Password, hash)

	fmt.Print(id, chech, err)

	return "alex", nil

}

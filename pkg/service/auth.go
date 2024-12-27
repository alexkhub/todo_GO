package service

import (
	"errors"
	"fmt"
	"strconv"
	todogo "todo"
	"todo/pkg/repository"

	jwt "github.com/dgrijalva/jwt-go"
)


type AuthService struct{
	repos repository.Authorization
	jwt_service JWTManager

}

func NewAuthService(repos repository.Authorization, jwt_service JWTManager ) *AuthService{
	return &AuthService{
		repos:  repos,
		jwt_service: jwt_service, 
	
	}
}

func (s *AuthService) CreateUser(user todogo.RegisterUser) (int, error){
	user.Password, _ = HashPassword(user.Password)
	
	return s.repos.CreateUser(user)
}


func (s *AuthService) LoginUser(user todogo.LoginUser) (todogo.JWTToken, error){
	
	id, hash, err := s.repos.GetUser(user)
	
	if err!= nil{
		return todogo.JWTToken{}, err
	}
	 
	if !CheckPasswordHash(user.Password, hash) {
		
		return todogo.JWTToken{}, errors.New("Password error")
	}
	
	access, err := s.jwt_service.CreateJwtAccess(strconv.Itoa(id))
	
	if err != nil{
		return todogo.JWTToken{}, errors.New("JWT error " + err.Error())	
	}

	
	refresh, err := s.jwt_service.CreateJwtRefresh(strconv.Itoa(id))
	
	if err != nil{
		return todogo.JWTToken{}, errors.New("JWT error " + err.Error())	
	}

	err = s.repos.CreateNewRefresh(HashUserId(id), refresh)

	if err != nil{
		return todogo.JWTToken{}, errors.New("JWT error " + err.Error())
	}


	return todogo.JWTToken{Access: access, Refresh: refresh}, nil

}

func (s *AuthService) RefreshJWT(refresh_token string) (todogo.JWTToken, error){

	token, err := jwt.Parse(refresh_token, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Parse error")
		}

		return []byte(signingKey2), nil
	})
	if err!= nil{
		return todogo.JWTToken{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return todogo.JWTToken{}, fmt.Errorf("error get user claims from token")
	}

	id := claims["sub"].(string)
	fmt.Print(id)
	

	
	access, err := s.jwt_service.CreateJwtAccess(id)
	
	if err != nil{
		return todogo.JWTToken{}, errors.New("JWT error " + err.Error())	
	}
	refresh, err := s.jwt_service.CreateJwtRefresh(id)
	
	if err != nil{
		return todogo.JWTToken{}, errors.New("JWT error " + err.Error())	
	}
	
    id2, _  := strconv.Atoi(id)

	err = s.repos.CreateNewRefresh(HashUserId(id2), refresh)

	if err != nil{
		return todogo.JWTToken{}, errors.New("JWT error " + err.Error())
	}


	return todogo.JWTToken{Access: access, Refresh: refresh}, nil


}
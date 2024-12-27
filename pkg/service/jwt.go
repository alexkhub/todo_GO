package service

import (
	"fmt"
	
	"time"


	jwt "github.com/dgrijalva/jwt-go"
)


type Manager struct{
	signingKey string
	signingKey2 string
}

func NewManager(signingKey string, signingKey2 string) *Manager{
	return &Manager{signingKey: signingKey, signingKey2: signingKey}
}

func (m *Manager) CreateJwtAccess(user_id string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		Subject: user_id,
	})
	return token.SignedString([]byte(m.signingKey))
}


func (m *Manager) CreateJwtRefresh(user_id string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		Subject: user_id,
	})
	return token.SignedString([]byte(m.signingKey2))
} 




func (m *Manager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
package service

import (
	"math"

	"golang.org/x/crypto/bcrypt"
)

const (
    salt = "fdgbgfd1232@$lv"
)


func HashPassword(password string) (string, error) {
    password = password + salt
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}
 
func CheckPasswordHash(password, hash string) bool {
    password = password + salt
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func HashUserId(user_id int)(float64){
    return math.Log10(float64(user_id))

}
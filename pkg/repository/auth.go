package repository

import (
	"fmt"

	todogo "todo"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres{
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres)CreateUser(user todogo.RegisterUser)(int, error){
	var id int
	query := fmt.Sprintf("Insert into %s (username, email, hash_password) values ($1, $2, $3) returning id", UserTable)

	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	

	if err:= row.Scan(&id); err != nil{
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser (user todogo.LoginUser)(int,  string,error){
	var id int
	var hash_password string

	query := fmt.Sprintf("select id, hash_password from %s where username = ($1) limit 1 ", UserTable)

	row := r.db.QueryRow(query, user.Username)

	if err:=row.Scan(&id, &hash_password); err != nil{
		return 0, "", err
	}

	return id, hash_password, nil
}

func (r *AuthPostgres) CreateNewRefresh(hash_id float64, refresh_token string)(error){
	var id int
	query := fmt.Sprintf("delete from %s where user_id = ($1) ", JWTTable)
	_, err := r.db.Exec(query, hash_id)
	if err != nil{
		return err
	}
	query = fmt.Sprintf("Insert into %s (user_id, refresh_token) values ($1, $2) returning id", JWTTable)

	row := r.db.QueryRow(query, hash_id, refresh_token)
	if err:= row.Scan(&id); err != nil{
		return err
	}
	return  nil

}
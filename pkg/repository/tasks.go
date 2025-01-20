package repository

import (
	"fmt"
	"strings"
	todogo "todo"

	"github.com/jmoiron/sqlx"
)


type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres{
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) CreateTask(user_id int, task todogo.CreateTask)(int, error){
	var id int
	query := fmt.Sprintf("Insert into %s (user_id, title, description) values ($1, $2, $3) returning id", TaskTable)

	row := r.db.QueryRow(query, user_id, task.Title, task.Description)
	

	if err:= row.Scan(&id); err != nil{
		return 0, err
	}
	return id, nil
}
func (r *TaskPostgres) TaskList(user_id int)([]todogo.ListTask, error){
	var lists[]todogo.ListTask
	query := fmt.Sprintf("SELECT id, title, description, status, datetime_create FROM %s WHERE user_id=$1;", TaskTable)
	err := r.db.Select(&lists, query, user_id)
	if err != nil{
		return lists, err
	}
	return lists, nil

}

func (r *TaskPostgres) TaskDetail(user_id int, task_id int)(todogo.ListTask, error){
	var task todogo.ListTask
	query := fmt.Sprintf("SELECT id, title, description, status, datetime_create FROM %s WHERE user_id=$1 and id=$2;", TaskTable)
	err := r.db.Get(&task, query, user_id, task_id)
	if err != nil{
		return task, err
	}

	return task, nil
}
	
func (r *TaskPostgres) TaskDelete(user_id int, task_id int)( error){
	query := fmt.Sprintf("Delete  FROM %s WHERE user_id=$1 and id=$2;", TaskTable)
	_, err := r.db.Exec(query, user_id, task_id)
	
	return err
}

func (r *TaskPostgres) TaskUpdate(user_id int, task_id int, input todogo.UpdateTask)(todogo.ListTask, error){
	var task todogo.ListTask
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	

	if input.Title != ""{
		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
		args = append(args, input.Title)
		argId ++
		
	}
	if input.Description != ""{
		setValue = append(setValue, fmt.Sprintf("description=$%d", argId))
		args = append(args, input.Description)
		argId ++
		
	}
	if input.Status != ""{
		setValue = append(setValue, fmt.Sprintf("status=$%d", argId))
		args = append(args, input.Status)
		argId ++
		
	}
	setQuery := strings.Join(setValue, ", ")

	query := fmt.Sprintf("Update %s set %s WHERE user_id=$%d and id=$%d;", TaskTable, setQuery, argId, argId+1)
	fmt.Print(query)
	args = append(args, user_id, task_id)

	_, err := r.db.Exec(query, args... )
	if err != nil{
		return task, err
	} 

	query = fmt.Sprintf("SELECT id, title, description, status, datetime_create FROM %s WHERE user_id=$1 and id=$2;", TaskTable)
	err = r.db.Get(&task, query, user_id, task_id)
	if err != nil{
		return task, err
	}


	return task, nil
}
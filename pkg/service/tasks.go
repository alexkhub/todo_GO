package service

import (
	todogo "todo"
	"todo/pkg/repository"
)

type TaskService struct{
	repos repository.Task
}

func NewTaskService(repos repository.Task) *TaskService{
	return &TaskService{
		repos: repos,
	}
}

func (t *TaskService) CreateTask(user_id int, task todogo.CreateTask)(int, error){
	return t.repos.CreateTask(user_id, task)
}

func (t *TaskService) TaskList(user_id int)([]todogo.ListTask, error){
	return t.repos.TaskList(user_id)
}

func (t *TaskService) TaskDetail(user_id int, task_id int)(todogo.ListTask, error){
	return t.repos.TaskDetail(user_id, task_id)
}

func (t *TaskService) TaskDelete(user_id int, task_id int)(error){
	return t.repos.TaskDelete(user_id, task_id)
}

func (t *TaskService) TaskUpdate(user_id int, task_id int, input todogo.UpdateTask)(todogo.ListTask, error){
	return t.repos.TaskUpdate(user_id, task_id, input)
}
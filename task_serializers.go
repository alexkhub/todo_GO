package todogo


type Task struct{
	Id int `json: "-"`
	User int `json: "user"  binding:"required"`
	Title int `json: "title"  binding:"required"`
	Description int `json: "description"  binding:"required"`
	Status int `json: "status"  binding:"required"`
	

}
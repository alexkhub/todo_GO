package todogo

type Task struct {
	Id          int `json:"id"`
	User        int `json:"user"  binding:"required" valid:"-"`
	Title      	string `json:"title"  binding:"required" valid:"-"`
	Description string `json:"description"  binding:"required" valid:"-"`
	Status      string `json:"status"  binding:"required" valid:"-"`
	DatetimeCreate string `json:"datetime_create"  binding:"required" valid:"-"`
}

type CreateTask struct{
	Title       string `json:"title" binding:"required" valid:"-"`
	Description string `json:"description" binding:"required" valid:"-"`
}

type ListTask struct {
	Id          int `json:"id" db:"id"`
	Title      	string `json:"title" binding:"required" valid:"-" db:"title"`
	Description string `json:"description"  binding:"required" valid:"-" db:"description"`
	Status      string `json:"status"  binding:"required" valid:"-"  db:"status"`
	DatetimeCreate string `json:"datetime_create"  binding:"required" valid:"-"  db:"datetime_create"`
}

type UpdateTask struct {
	Title      	string `json:"title" valid:"-" db:"title"`
	Description string `json:"description" valid:"-" db:"description"`
	Status      string `json:"status" valid:"-" db:"status"`
}


package handlers


import (
	
	"github.com/gin-gonic/gin"
	"database/sql"
)



type Handler struct{
    DB  *sql.DB
}

func (h *Handler) InitRouter() * gin.Engine{
	router:= gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/rigister", h.Register)
			auth.POST("/login", h.Login)
			auth.GET("/all_user", h.UserList)
		}
		task := api.Group("/task")
		{
			task.GET("/", h.TaskList)
			task.GET("task_detail/:id", h.TaskDetail)
			task.POST("new_task/", h.TaskCreate)
			task.PATCH("task_detail/:id", h.TaskUpdate)
			task.DELETE("task_detail/:id", h.TaskDelete)
			
		}
	}
	return router
}
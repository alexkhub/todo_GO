package handlers

import (
	
	"todo/pkg/service"

	"github.com/gin-gonic/gin"
	
	
)

type Handler struct{
    servies *service.Service
	
}


func NewHandler(services *service.Service) *Handler{
	return &Handler{servies: services}
}

func (h *Handler) InitRouter() * gin.Engine{
	router:= gin.New()
	router.GET("/", h.MainPage)
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			
			auth.POST("/rigister", h.Register)
			auth.POST("/login", h.Login)
			
			
		}
		task := api.Group("/task", )
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
package handlers

import (
	
	"todo/pkg/service"

	"github.com/gin-gonic/gin"
	
	
)

type Handler struct{
    servies *service.Service
	auth service.JWTManager
	
	
}


func NewHandler(services *service.Service, auth service.JWTManager) *Handler{
	return &Handler{servies: services, auth: auth}
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
			auth.POST("/refresh", h.RefreshToken )
			
			
		}
		task := api.Group("/task",  h.parseAuthHeader)
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
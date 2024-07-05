package handler

import (
	_ "timeTracker/docs"
	"timeTracker/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.DELETE("/:id", h.deleteUser)
			users.PUT("/:id", h.updateUser)
			users.GET("/", h.getPeopleByFilters)

			tasks := users.Group(":id/tasks")
			{
				tasks.POST("/", h.createTask)
			}
		}
		tasks := api.Group("/tasks")
		{
			tasks.GET("/:id", h.getTasksByLaborCosts)
			tasks.PUT("/start/:id", h.startTask)
			tasks.PUT("/end/:id", h.endTask)
			tasks.DELETE("/:id", h.deleteTask)
		}
	}

	return router
}

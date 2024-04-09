package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/morf1lo/todo/internal/service"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.SetTrustedProxies(nil)

	router.Use(cors.Default())

	auth := router.Group("/api/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
	}

	user := router.Group("/api/users")
	{
		user.GET("/", h.authMiddleware, h.getUser)
		user.PUT("/", h.authMiddleware, h.updateUsername)
	}

	todo := router.Group("/api/todos")
	{
		todo.POST("/create", h.authMiddleware, h.createTodo)
		todo.GET("/", h.authMiddleware, h.getTodos)
		todo.PUT("/:id", h.authMiddleware, h.updateTodo)
		todo.DELETE("/:id", h.authMiddleware, h.deleteTodo)
		todo.GET("/completed", h.authMiddleware, h.getCompletedTodos)
		todo.GET("/important", h.authMiddleware, h.getImportantTodos)
		todo.GET("/uncompleted", h.authMiddleware, h.getUncompletedTodos)
	}

	return router
}

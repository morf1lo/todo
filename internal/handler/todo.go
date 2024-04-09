package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morf1lo/todo/internal/models"
	"github.com/morf1lo/todo/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) createTodo(c *gin.Context) {
	user := auth.GetUserFromRequest(c)

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	todo.UserID = user.ID

	if err := h.services.Todo.Create(c.Request.Context(), todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) getTodos(c *gin.Context) {
	user := auth.GetUserFromRequest(c)

	todos, err := h.services.Todo.GetAll(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "data": todos})
}

func (h *Handler) updateTodo(c *gin.Context) {
	user := auth.GetUserFromRequest(c)

	var options models.TodoUpdateOptions
	if err := c.ShouldBindJSON(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoIDString := c.Param("id")
	todoID, err := primitive.ObjectIDFromHex(todoIDString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Todo.Update(c.Request.Context(), todoID, user.ID, options); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) deleteTodo(c *gin.Context) {
	user := auth.GetUserFromRequest(c)

	todoIDString := c.Param("id")
	todoID, err := primitive.ObjectIDFromHex(todoIDString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Todo.Delete(c.Request.Context(), todoID, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) getCompletedTodos(c *gin.Context) {
	user := auth.GetUserFromRequest(c)

	todos, err := h.services.Todo.GetCompletedTodos(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "data": todos})
}

func (h *Handler) getImportantTodos(c *gin.Context) {
	user := auth.GetUserFromRequest(c)

	todos, err := h.services.Todo.GetImportantTodos(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "data": todos})
}

func (h *Handler) getUncompletedTodos(c *gin.Context) {
	user := auth.GetUserFromRequest(c)

	todos, err := h.services.Todo.GetUncompletedTodos(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "data": todos})
}

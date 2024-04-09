package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/morf1lo/todo/internal/models"
	"github.com/morf1lo/todo/pkg/auth"
)

func (h *Handler) getUser(c *gin.Context) {
	user := auth.GetUserFromRequest(c)

	dbUser, err := h.services.User.GetById(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "data": dbUser})
}

func (h *Handler) updateUsername(c *gin.Context) {
	user := auth.GetUserFromRequest(c)

	var newUsername models.User
	if err := c.ShouldBindJSON(&newUsername); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUsername.Username = strings.TrimSpace(newUsername.Username)
	if newUsername.Username == "" || len(newUsername.Username) < 3 || len(newUsername.Username) > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username must be at least 3 and no more than 12"})
		return
	}

	if err := h.services.User.UpdateUsername(c.Request.Context(), user.ID, newUsername.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

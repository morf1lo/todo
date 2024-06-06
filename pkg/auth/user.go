package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/morf1lo/todo/internal/model"
)

func GetUserFromRequest(c *gin.Context) model.User {
	claims, _ := c.Get("user")

	user, ok := claims.(model.User)
	if !ok {
		return model.User{}
	}

	return user
}

package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) authMiddleware(c *gin.Context) {
	tokenCookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not authorized"})
		c.Abort()
		return
	}
	
	parsedToken, err := jwt.Parse(tokenCookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot parse authentication token"})
		c.Abort()
		return
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication token is not valid"})
		c.Abort()
		return
	}

	idString := claims["id"].(string)
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token claims"})
		c.Abort()
		return
	}

	user, err := h.services.User.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Set("user", *user)

	c.Next()
}

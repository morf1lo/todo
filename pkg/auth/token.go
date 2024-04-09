package auth

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func generateToken(userID string) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	token, err := jwt.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func CreateSendToken(c *gin.Context, userID string) error {
	jwt, err := generateToken(userID)
	if err != nil {
		return err
	}

	c.SetCookie("jwt", jwt, 3600 * 24 * 7, "/", "localhost", true, true)
	return nil
}

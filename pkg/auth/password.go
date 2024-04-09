package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 15)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func VerifyPassword(hashed []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, password)
	return err == nil
}

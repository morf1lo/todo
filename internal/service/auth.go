package service

import (
	"context"

	"github.com/morf1lo/todo/internal/model"
	"github.com/morf1lo/todo/internal/repository"
	"github.com/morf1lo/todo/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Create(ctx context.Context, user model.User) (string, error) {
	user.ID = primitive.NewObjectID()
	if err := s.repo.Authorization.CreateUser(ctx, user); err != nil {
		return "", err
	}
	
	return user.ID.Hex(), nil
}

func (s *AuthService) SignIn(ctx context.Context, user model.User) (string, error) {
	var existingUser *model.User
	if err := s.repo.User.FindByUsername(ctx, user.Username).Decode(&existingUser); err != nil {
		return "", err
	}
	if existingUser == nil {
		return "", errUserNotFound
	}

	if ok := auth.VerifyPassword([]byte(existingUser.Password), []byte(user.Password)); !ok {
		return "", errInvalidCredentials
	}

	return existingUser.ID.Hex(), nil
}

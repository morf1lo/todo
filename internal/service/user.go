package service

import (
	"context"

	"github.com/morf1lo/todo/internal/models"
	"github.com/morf1lo/todo/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetById(ctx context.Context, userID primitive.ObjectID) (*models.User, error) {
	var user models.User
	if err := s.repo.User.FindById(ctx, userID).Decode(&user); err != nil {
		return nil, err
	}

	modifiedUser := models.User{
		ID: user.ID,
		Username: user.Username,
	}

	return &modifiedUser, nil
}

func (s *UserService) UpdateUsername(ctx context.Context, userID primitive.ObjectID, newUsername string) error {
	if err := s.repo.User.UpdateUsername(ctx, userID, newUsername); err != nil {
		return nil
	}
	return nil
}

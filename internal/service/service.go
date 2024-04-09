package service

import (
	"context"

	"github.com/morf1lo/todo/internal/models"
	"github.com/morf1lo/todo/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	Create(ctx context.Context, user models.User) (string, error)
	SignIn(ctx context.Context, user models.User) (string, error)
}

type User interface {
	GetById(ctx context.Context, userID primitive.ObjectID) (*models.User, error)
	UpdateUsername(ctx context.Context, userID primitive.ObjectID, newUsername string) error
}

type Todo interface {
	Create(ctx context.Context, todo models.Todo) error
	GetAll(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error)
	Update(ctx context.Context, todoID primitive.ObjectID, userID primitive.ObjectID, options models.TodoUpdateOptions) error
	Delete(ctx context.Context, todoID primitive.ObjectID, userID primitive.ObjectID) error
	GetCompletedTodos(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error)
	GetImportantTodos(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error)
	GetUncompletedTodos(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error)
}

type Service struct {
	Authorization
	User
	Todo
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		User: NewUserService(repo),
		Todo: NewTodoService(repo),
	}
}

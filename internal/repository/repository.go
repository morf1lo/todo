package repository

import (
	"context"

	"github.com/morf1lo/todo/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(ctx context.Context, user model.User) error
}

type User interface {
	FindById(ctx context.Context, userID primitive.ObjectID) *mongo.SingleResult
	FindByUsername(ctx context.Context, username string) *mongo.SingleResult 
	UpdateUsername(ctx context.Context, userID primitive.ObjectID, newUsername string) error
}

type Todo interface {
	Create(ctx context.Context, todo model.Todo) error
	FindAll(ctx context.Context, userID primitive.ObjectID) ([]*model.Todo, error)
	Update(ctx context.Context, todoID primitive.ObjectID, userID primitive.ObjectID, options model.TodoUpdateOptions) error
	Delete(ctx context.Context, todoID primitive.ObjectID, userID primitive.ObjectID) error
	FindCompletedTodos(ctx context.Context, userID primitive.ObjectID) ([]*model.Todo, error)
	FindImportantTodos(ctx context.Context, userID primitive.ObjectID) ([]*model.Todo, error)
	FindUncompletedTodos(ctx context.Context, userID primitive.ObjectID) ([]*model.Todo, error)
}

type Repository struct {
	Authorization
	User
	Todo
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		User: NewUserRepository(db),
		Todo: NewTodoRepository(db),
	}
}
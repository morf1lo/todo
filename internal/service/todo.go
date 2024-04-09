package service

import (
	"context"
	"time"

	"github.com/morf1lo/todo/internal/models"
	"github.com/morf1lo/todo/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService struct {
	repo *repository.Repository
}

func NewTodoService(repo *repository.Repository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(ctx context.Context, todo models.Todo) error {
	timestamp := time.Now()
	todo.CreatedAt = timestamp.Format("2006-01-02 15:04:05")
	todo.ID = primitive.NewObjectID()

	if err := s.repo.Todo.Create(ctx, todo); err != nil {
		return err
	}
	return nil
}

func (s *TodoService) GetAll(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error) {
	todos, err := s.repo.Todo.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) Update(ctx context.Context, todoID primitive.ObjectID, userID primitive.ObjectID, options models.TodoUpdateOptions) error {
	if err := s.repo.Todo.Update(ctx, todoID, userID, options); err != nil {
		return err
	}
	return nil
}

func (s *TodoService) Delete(ctx context.Context, todoID primitive.ObjectID, userID primitive.ObjectID) error {
	if err := s.repo.Todo.Delete(ctx, todoID, userID); err != nil {
		return nil
	}
	return nil
}

func (s *TodoService) GetCompletedTodos(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error) {
	todos, err := s.repo.Todo.FindCompletedTodos(ctx, userID)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetImportantTodos(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error) {
	todos, err := s.repo.Todo.FindImportantTodos(ctx, userID)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetUncompletedTodos(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error) {
	todos, err := s.repo.Todo.FindImportantTodos(ctx, userID)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

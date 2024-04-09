package repository

import (
	"context"

	"github.com/morf1lo/todo/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	db *mongo.Database
}

func NewTodoRepository(db *mongo.Database) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(ctx context.Context, todo models.Todo) error {
	_, err := r.db.Collection("todos").InsertOne(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) FindAll(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error) {
	cursor, err := r.db.Collection("todos").Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []*models.Todo
	for cursor.Next(ctx) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepository) Update(ctx context.Context, todoID primitive.ObjectID, userID primitive.ObjectID, options models.TodoUpdateOptions) error {
	result := r.db.Collection("todos").FindOne(ctx, bson.M{"_id": todoID, "userId": userID})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return errTodoNotFound
		}
		return err
	}
	
	query := options.GetUpdateQuery()

	_, err := r.db.Collection("todos").UpdateByID(ctx, todoID, query)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) Delete(ctx context.Context, todoID primitive.ObjectID, userID primitive.ObjectID) error {
	_, err := r.db.Collection("todos").DeleteOne(ctx, bson.M{"_id": todoID, "userId": userID})
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) FindCompletedTodos(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error) {
	cursor, err := r.db.Collection("todos").Find(ctx, bson.M{"userId": userID, "completed": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []*models.Todo
	for cursor.Next(ctx) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepository) FindImportantTodos(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error) {
	cursor, err := r.db.Collection("todos").Find(ctx, bson.M{"userId": userID, "important": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []*models.Todo
	for cursor.Next(ctx) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepository) FindUncompletedTodos(ctx context.Context, userID primitive.ObjectID) ([]*models.Todo, error) {
	cursor, err := r.db.Collection("todos").Find(ctx, bson.M{"userId": userID, "completed": false})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []*models.Todo
	for cursor.Next(ctx) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

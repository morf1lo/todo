package repository

import (
	"context"

	"github.com/morf1lo/todo/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	db *mongo.Database
}

func NewAuthRepository(db *mongo.Database) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user model.User) error {
	result := r.db.Collection("users").FindOne(ctx, bson.M{"username": user.Username})
	if err := result.Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	} else {
		return errUsernameIsAlreadyTaken
	}

	_, err := r.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

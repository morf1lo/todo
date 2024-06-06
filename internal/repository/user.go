package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(ctx context.Context, userID primitive.ObjectID) *mongo.SingleResult {
	result := r.db.Collection("users").FindOne(ctx, bson.M{"_id": userID})
	return result
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) *mongo.SingleResult {
	result := r.db.Collection("users").FindOne(ctx, bson.M{"username": username})
	return result
}

func (r *UserRepository) UpdateUsername(ctx context.Context, userID primitive.ObjectID, newUsername string) error {
	result := r.db.Collection("users").FindOne(ctx, bson.M{"username": newUsername})
	if err := result.Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	} else {
		return errUsernameIsAlreadyTaken
	}

	_, err := r.db.Collection("users").UpdateByID(ctx, userID, bson.M{"$set": bson.M{"username": newUsername}})
	if err != nil {
		return err
	}

	return nil
}
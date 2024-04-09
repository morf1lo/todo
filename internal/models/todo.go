package models

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Body      string             `json:"body" binding:"required,max=80"`
	Completed bool               `json:"completed"`
	Important bool               `json:"important"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
}

type TodoUpdateOptions struct {
	Body      *string `json:"body"`
	Completed *bool   `json:"completed"`
	Important *bool   `json:"important"`
}

func (t *TodoUpdateOptions) GetUpdateQuery() bson.M {
	query := bson.M{"$set": bson.M{}}

	if t.Body != nil && len(strings.TrimSpace(*t.Body)) <= 80 {
		query["$set"].(bson.M)["body"] = strings.TrimSpace(*t.Body)
	}

	if t.Completed != nil {
		query["$set"].(bson.M)["completed"] = t.Completed
	}

	if t.Important != nil {
		query["$set"].(bson.M)["important"] = t.Important
	}

	return query
}

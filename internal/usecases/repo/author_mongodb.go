package repo

import (
	"context"
	"fmt"
	"go-blog-ca/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorRepo struct {
	collection *mongo.Collection // or *mongo.Database
}

func NewAuthorRepo(db *mongo.Database) *AuthorRepo {
	return &AuthorRepo{
		collection: db.Collection("authors"),
	}
}

func (r *AuthorRepo) Create(ctx context.Context, author entity.Author) (string, error) {
	result, err := r.collection.InsertOne(ctx, author)
	if err != nil {
		return "", fmt.Errorf("failed to create author due to error: %v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("failed to convert oid to hex (oid):%s", oid)
}

func (r *AuthorRepo) FindByID(ctx context.Context, id string) (entity.Author, error) {
	panic("Implement me")
}

func (r *AuthorRepo) FindAll(ctx context.Context) ([]entity.Author, error) {
	panic("Implement me")
}

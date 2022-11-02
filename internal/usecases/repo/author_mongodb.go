package repo

import (
	"context"
	"go-blog-ca/internal/domain/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorRepo struct {
	*mongo.Collection // or *mongo.Database
}

func NewAuthorRepo(db *mongo.Database) *AuthorRepo {
	return &AuthorRepo{
		db.Collection("authors"),
	}
}

func (r *AuthorRepo) Create(ctx context.Context, author entity.Author) (string, error) {
	panic("Implement me")
}

func (r *AuthorRepo) FindByID(ctx context.Context, id string) (entity.Author, error) {
	panic("Implement me")
}

func (r *AuthorRepo) FindAll(ctx context.Context) ([]entity.Author, error) {
	panic("Implement me")
}

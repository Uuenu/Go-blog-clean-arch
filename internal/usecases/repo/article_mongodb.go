package repo

import (
	"context"
	"go-blog-ca/internal/domain/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleRepo struct {
	*mongo.Collection // or *mongo.Database
}

func NewArticleRepo(db *mongo.Database) *ArticleRepo {
	return &ArticleRepo{
		db.Collection("articles"),
	}
}

func (r *ArticleRepo) Create(context.Context, entity.Article) (string, error) {
	panic("Implement me")
}

func (r *ArticleRepo) FindById(ctx context.Context, id string) (entity.Article, error) {
	panic("Implement me")
}

func (r *ArticleRepo) FindAll(ctx context.Context) ([]entity.Article, error) {
	panic("Implement me")
}

func (r *ArticleRepo) Delete(ctx context.Context, id string, aid string) error {
	panic("Implement me")
}

package service

import (
	"context"
	"go-blog-ca/internal/domain/entity"
)

type ArticleStorage interface {
	Create(article entity.Article) error
	GetOne(ID string) (entity.Article, error)
	GetAll() ([]entity.Article, error)
}

type articleService struct {
	Storage ArticleStorage
}

func NewArticleService(storage ArticleStorage) *articleService {
	return &articleService{
		Storage: storage,
	}
}

func (s articleService) GetByID(ctx context.Context, id string) (entity.Article, error) {
	return s.Storage.GetOne(id)
}

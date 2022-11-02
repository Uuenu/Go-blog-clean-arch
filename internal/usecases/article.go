package usecases

import (
	"context"
	"go-blog-ca/internal/domain/entity"
)

type ArticleUseCases struct {
	repo ArticleRepo
}

func NewArticleUseCase(r ArticleRepo) *ArticleUseCases {
	return &ArticleUseCases{
		repo: r,
	}
}

func (uc *ArticleUseCases) Create(context.Context, entity.Article) (string, error) {
	panic("Implement me")
}

func (uc *ArticleUseCases) GetByID(ctx context.Context, id string) (entity.Article, error) {
	panic("Implement me")
}

func (uc *ArticleUseCases) GetAll(ctx context.Context) ([]entity.Article, error) {
	panic("Implement me")
}

func (uc *ArticleUseCases) Delete(ctx context.Context, id string, aid string) error {
	panic("Implement me")
}

package usecases

import (
	"context"
	"fmt"
	"go-blog-ca/internal/domain/entity"
	"go-blog-ca/pkg/apperrors"
)

type ArticleUseCases struct {
	repo    ArticleRepo
	author  Author
	session Session
}

func NewArticleUseCase(r ArticleRepo, a Author, s Session) *ArticleUseCases {
	return &ArticleUseCases{
		repo:    r,
		author:  a,
		session: s,
	}
}

func (uc *ArticleUseCases) Create(ctx context.Context, article entity.Article) (string, error) {
	//TODO change entity.Article on DTO or somthing else
	aid, err := uc.repo.Create(ctx, article)
	if err != nil {
		return "", fmt.Errorf("ArticleUseCase - Create - uc.repo.Create. error: %v", err)
	}

	return aid, nil
}

func (uc *ArticleUseCases) GetByID(ctx context.Context, id string) (entity.Article, error) {
	article, err := uc.repo.FindById(ctx, id)
	if err != nil {
		return entity.Article{}, fmt.Errorf("ArticleUseCase - GetByID - uc.repo.FindByID. error: %v", err)
	}

	return article, nil
}

func (uc *ArticleUseCases) GetAll(ctx context.Context) ([]entity.Article, error) {
	articles, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("ArticleUseCase - GetAll - uc.repo.FindAll. error: %v", err)
	}

	return articles, nil
}

func (uc *ArticleUseCases) Delete(ctx context.Context, id string, sid string) error {

	sess, err := uc.session.GetByID(ctx, sid)
	if err != nil {
		return fmt.Errorf("ArticleUseCase - Delete - uc.session.GetByID: %v", err)
	}

	article, err := uc.repo.FindById(ctx, id)
	if err != nil {
		return fmt.Errorf("ArticleUseCase - Delete - uc.repo.FindByID: %v", err)
	}

	if sess.AuthorID == article.AuthorID {
		return uc.repo.Delete(ctx, id, article.AuthorID)
	}

	return fmt.Errorf("ArticleUseCase - Delete: %v", apperrors.ErrArticleAccessDenied)
}

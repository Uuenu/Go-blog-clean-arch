package usecases

import (
	"context"
	"go-blog-ca/internal/domain/entity"
)

type AuthorUseCase struct {
	repo AuthorRepo
}

func NewAuthorUseCase(r AuthorRepo) *AuthorUseCase {
	return &AuthorUseCase{
		repo: r,
	}
}

func (uc *AuthorUseCase) Create(ctx context.Context, author entity.Author) (string, error) {
	panic("Implement me")
}

func (uc *AuthorUseCase) GetByID(ctx context.Context, id string) (entity.Author, error) {
	panic("Implement me")
}

func (uc *AuthorUseCase) GetAll(ctx context.Context) ([]entity.Author, error) {
	panic("Implement me")
}
